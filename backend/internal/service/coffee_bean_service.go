package service

import (
	"errors"
	"fmt"

	"my-coffee-log/internal/model"
	"my-coffee-log/internal/repository"
	"my-coffee-log/internal/utils"
)

type CoffeeBeanService struct {
	beanRepo *repository.CoffeeBeanRepository
}

func NewCoffeeBeanService(beanRepo *repository.CoffeeBeanRepository) *CoffeeBeanService {
	return &CoffeeBeanService{beanRepo: beanRepo}
}

type CreateCoffeeBeanRequest struct {
	Name             string `json:"name" binding:"required"`
	Origin           string `json:"origin"`
	ProcessingMethod string `json:"processing_method"`
	RoastLevel       string `json:"roast_level"`
	Roaster          string `json:"roaster"`
	ImageURL         string `json:"image_url"`
}

type UpdateCoffeeBeanRequest struct {
	Name             *string `json:"name"`
	Origin           *string `json:"origin"`
	ProcessingMethod *string `json:"processing_method"`
	RoastLevel       *string `json:"roast_level"`
	Roaster          *string `json:"roaster"`
	ImageURL         *string `json:"image_url"`
}

func (s *CoffeeBeanService) Create(userID uint, req CreateCoffeeBeanRequest) (*model.CoffeeBean, error) {
	if len([]rune(req.Name)) > 255 {
		return nil, errors.New("bean name is too long")
	}

	existing, _ := s.beanRepo.FindByName(userID, req.Name)
	if existing != nil {
		return nil, fmt.Errorf("bean '%s' already exists", req.Name)
	}

	bean := &model.CoffeeBean{
		UserID:           userID,
		Name:             req.Name,
		Origin:           req.Origin,
		ProcessingMethod: req.ProcessingMethod,
		RoastLevel:       req.RoastLevel,
		Roaster:          req.Roaster,
		ImageURL:         req.ImageURL,
	}

	if err := s.beanRepo.Create(bean); err != nil {
		return nil, err
	}

	return bean, nil
}

func (s *CoffeeBeanService) GetByID(id, userID uint) (*model.CoffeeBean, error) {
	bean, err := s.beanRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("coffee bean not found")
	}
	return bean, nil
}

func (s *CoffeeBeanService) GetList(userID uint, pagination *utils.Pagination, search string) ([]model.CoffeeBean, int64, error) {
	return s.beanRepo.FindList(userID, pagination, search)
}

func (s *CoffeeBeanService) Update(id, userID uint, req UpdateCoffeeBeanRequest) (*model.CoffeeBean, error) {
	bean, err := s.beanRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("coffee bean not found")
	}

	if req.Name != nil {
		if len([]rune(*req.Name)) > 255 {
			return nil, errors.New("bean name is too long")
		}
		existing, _ := s.beanRepo.FindByName(userID, *req.Name)
		if existing != nil && existing.ID != id {
			return nil, fmt.Errorf("bean '%s' already exists", *req.Name)
		}
		bean.Name = *req.Name
	}
	if req.Origin != nil {
		bean.Origin = *req.Origin
	}
	if req.ProcessingMethod != nil {
		bean.ProcessingMethod = *req.ProcessingMethod
	}
	if req.RoastLevel != nil {
		bean.RoastLevel = *req.RoastLevel
	}
	if req.Roaster != nil {
		bean.Roaster = *req.Roaster
	}
	if req.ImageURL != nil {
		bean.ImageURL = *req.ImageURL
	}

	if err := s.beanRepo.Update(bean); err != nil {
		return nil, err
	}

	return bean, nil
}

func (s *CoffeeBeanService) Delete(id, userID uint) error {
	return s.beanRepo.Delete(id, userID)
}

func (s *CoffeeBeanService) GetBeanList(userID uint) ([]model.CoffeeBean, error) {
	return s.beanRepo.FindBeanNames(userID)
}

// EnsureBeanForLog ensures a CoffeeBean record exists when referenced in a coffee log.
// If bean_id is provided, increment usage count. If bean info is provided inline, create a new bean.
func (s *CoffeeBeanService) EnsureBeanForLog(userID uint, beanID *uint, beanName string) (*uint, error) {
	if beanID != nil && *beanID > 0 {
		// Existing bean reference - increment usage count
		if err := s.beanRepo.IncrementUsageCount(*beanID, userID); err != nil {
			return nil, err
		}
		return beanID, nil
	}

	// If inline bean name provided, auto-create
	if beanName == "" {
		return nil, nil
	}

	existing, err := s.beanRepo.FindByName(userID, beanName)
	if err == nil && existing != nil {
		// Bean exists, increment usage
		_ = s.beanRepo.IncrementUsageCount(existing.ID, userID)
		id := existing.ID
		return &id, nil
	}

	// Create new bean
	bean := &model.CoffeeBean{
		UserID:     userID,
		Name:       beanName,
		UsageCount: 1,
	}
	if err := s.beanRepo.Create(bean); err != nil {
		return nil, err
	}
	id := bean.ID
	return &id, nil
}
