package service

import (
	"errors"
	"fmt"
	"time"

	"my-coffee-log/internal/model"
	"my-coffee-log/internal/repository"
	"my-coffee-log/internal/utils"
)

type CoffeeShopService struct {
	shopRepo      *repository.CoffeeShopRepository
	coffeeLogRepo *repository.CoffeeLogRepository
}

func NewCoffeeShopService(shopRepo *repository.CoffeeShopRepository, coffeeLogRepo *repository.CoffeeLogRepository) *CoffeeShopService {
	return &CoffeeShopService{shopRepo: shopRepo, coffeeLogRepo: coffeeLogRepo}
}

type CreateCoffeeShopRequest struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address"`
	Rating   int    `json:"rating"`
	ImageURL string `json:"image_url"`
}

type UpdateCoffeeShopRequest struct {
	Name     *string `json:"name"`
	Address  *string `json:"address"`
	Rating   *int    `json:"rating"`
	ImageURL *string `json:"image_url"`
}

func (s *CoffeeShopService) Create(userID uint, req CreateCoffeeShopRequest) (*model.CoffeeShop, error) {
	if len([]rune(req.Name)) > 255 {
		return nil, errors.New("shop name is too long")
	}
	if req.Rating < 0 || req.Rating > 5 {
		return nil, errors.New("rating must be between 0 and 5")
	}

	// Check if shop with same name already exists for this user
	existing, _ := s.shopRepo.FindByName(userID, req.Name)
	if existing != nil {
		return nil, fmt.Errorf("shop '%s' already exists", req.Name)
	}

	shop := &model.CoffeeShop{
		UserID:   userID,
		Name:     req.Name,
		Address:  req.Address,
		Rating:   req.Rating,
		ImageURL: req.ImageURL,
	}

	if err := s.shopRepo.Create(shop); err != nil {
		return nil, err
	}

	return shop, nil
}

func (s *CoffeeShopService) GetByID(id, userID uint) (*model.CoffeeShop, error) {
	shop, err := s.shopRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("coffee shop not found")
	}
	return shop, nil
}

func (s *CoffeeShopService) GetList(userID uint, pagination *utils.Pagination, search string) ([]model.CoffeeShop, int64, error) {
	return s.shopRepo.FindList(userID, pagination, search)
}

func (s *CoffeeShopService) Update(id, userID uint, req UpdateCoffeeShopRequest) (*model.CoffeeShop, error) {
	shop, err := s.shopRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("coffee shop not found")
	}

	if req.Name != nil {
		if len([]rune(*req.Name)) > 255 {
			return nil, errors.New("shop name is too long")
		}
		// Check name conflict
		existing, _ := s.shopRepo.FindByName(userID, *req.Name)
		if existing != nil && existing.ID != id {
			return nil, fmt.Errorf("shop '%s' already exists", *req.Name)
		}
		shop.Name = *req.Name
	}
	if req.Address != nil {
		shop.Address = *req.Address
	}
	if req.Rating != nil {
		if *req.Rating < 0 || *req.Rating > 5 {
			return nil, errors.New("rating must be between 0 and 5")
		}
		shop.Rating = *req.Rating
	}
	if req.ImageURL != nil {
		shop.ImageURL = *req.ImageURL
	}

	if err := s.shopRepo.Update(shop); err != nil {
		return nil, err
	}

	return shop, nil
}

func (s *CoffeeShopService) Delete(id, userID uint) error {
	return s.shopRepo.Delete(id, userID)
}

func (s *CoffeeShopService) GetShopNames(userID uint) ([]string, error) {
	return s.shopRepo.FindShopNames(userID)
}

// EnsureShopForLog ensures a CoffeeShop record exists when a coffee log is created with a shop_name.
// If the shop already exists, increment visit count. If not, create it.
func (s *CoffeeShopService) EnsureShopForLog(userID uint, shopName string) error {
	if shopName == "" || shopName == "Local Coffee Spot" {
		return nil
	}

	existing, err := s.shopRepo.FindByName(userID, shopName)
	if err == nil && existing != nil {
		// Shop exists, increment visit count
		return s.shopRepo.IncrementVisitCount(existing.ID, userID)
	}

	// Create new shop
	shop := &model.CoffeeShop{
		UserID:      userID,
		Name:        shopName,
		VisitCount:  1,
		LastVisitAt: &time.Time{},
	}
	now := time.Now()
	shop.LastVisitAt = &now

	return s.shopRepo.Create(shop)
}

// GetRelatedLogs fetches coffee logs that belong to this shop (by shop_name match)
func (s *CoffeeShopService) GetRelatedLogs(shopID, userID uint, pagination *utils.Pagination) ([]model.CoffeeLog, int64, error) {
	shop, err := s.shopRepo.FindByID(shopID, userID)
	if err != nil {
		return nil, 0, errors.New("coffee shop not found")
	}

	return s.coffeeLogRepo.FindByShopName(userID, shop.Name, pagination)
}
