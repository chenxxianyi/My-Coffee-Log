package service

import (
	"errors"
	"fmt"
	"time"

	"my-coffee-log/internal/model"
	"my-coffee-log/internal/repository"
	"my-coffee-log/internal/utils"
)

type CoffeeLogService struct {
	coffeeLogRepo *repository.CoffeeLogRepository
	flavorTagRepo *repository.FlavorTagRepository
	aiService     *AIService
}

func NewCoffeeLogService(coffeeLogRepo *repository.CoffeeLogRepository, flavorTagRepo *repository.FlavorTagRepository, aiService *AIService) *CoffeeLogService {
	return &CoffeeLogService{
		coffeeLogRepo: coffeeLogRepo,
		flavorTagRepo: flavorTagRepo,
		aiService:     aiService,
	}
}

type CreateCoffeeLogRequest struct {
	CoffeeName  string `json:"coffee_name"`
	CoffeeType  string `json:"coffee_type"`
	ShopName    string `json:"shop_name"`
	Location    string `json:"location"`
	ImageURL    string `json:"image_url"`
	DrinkDate   string `json:"drink_date"`
	Mood        string `json:"mood"`
	Notes       string `json:"notes"`
	Acidity     int    `json:"acidity"`
	Bitterness  int    `json:"bitterness"`
	Sweetness   int    `json:"sweetness"`
	Body        int    `json:"body"`
	Aroma       int    `json:"aroma"`
	Aftertaste  int    `json:"aftertaste"`
	FlavorTagIDs []uint `json:"flavor_tag_ids"`
}

type UpdateCoffeeLogRequest struct {
	CoffeeName  string `json:"coffee_name"`
	CoffeeType  string `json:"coffee_type"`
	ShopName    string `json:"shop_name"`
	Location    string `json:"location"`
	ImageURL    string `json:"image_url"`
	DrinkDate   string `json:"drink_date"`
	Mood        string `json:"mood"`
	Notes       string `json:"notes"`
	Acidity     *int   `json:"acidity"`
	Bitterness  *int   `json:"bitterness"`
	Sweetness   *int   `json:"sweetness"`
	Body        *int   `json:"body"`
	Aroma       *int   `json:"aroma"`
	Aftertaste  *int   `json:"aftertaste"`
	FlavorTagIDs []uint `json:"flavor_tag_ids"`
}

func (s *CoffeeLogService) validateFlavorScores(acidity, bitterness, sweetness, body, aroma, aftertaste int) error {
	dimensions := []struct {
		name  string
		value int
	}{
		{"acidity", acidity},
		{"bitterness", bitterness},
		{"sweetness", sweetness},
		{"body", body},
		{"aroma", aroma},
		{"aftertaste", aftertaste},
	}
	for _, d := range dimensions {
		if d.value < 0 || d.value > 5 {
			return fmt.Errorf("%s must be between 0 and 5", d.name)
		}
	}
	return nil
}

func (s *CoffeeLogService) Create(userID uint, req CreateCoffeeLogRequest) (*model.CoffeeLog, error) {
	if err := s.validateFlavorScores(req.Acidity, req.Bitterness, req.Sweetness, req.Body, req.Aroma, req.Aftertaste); err != nil {
		return nil, err
	}

	if len(req.FlavorTagIDs) > 0 {
		tags, err := s.flavorTagRepo.FindByIDs(req.FlavorTagIDs)
		if err != nil || len(tags) != len(req.FlavorTagIDs) {
			return nil, errors.New("one or more flavor tags not found")
		}
	}

	var drinkDate *time.Time
	if req.DrinkDate != "" {
		parsed, err := time.Parse("2006-01-02", req.DrinkDate)
		if err != nil {
			return nil, errors.New("invalid drink_date format, expected YYYY-MM-DD")
		}
		drinkDate = &parsed
	}

	tagNames := []string{}
	if len(req.FlavorTagIDs) > 0 {
		tags, _ := s.flavorTagRepo.FindByIDs(req.FlavorTagIDs)
		for _, t := range tags {
			tagNames = append(tagNames, t.Name)
		}
	}

	aiSummary := s.aiService.GenerateMockSummary(req.CoffeeType, tagNames, req.Acidity, req.Bitterness, req.Sweetness, req.Body, req.Aroma, req.Aftertaste)

	log := &model.CoffeeLog{
		UserID:     userID,
		CoffeeName: req.CoffeeName,
		CoffeeType: req.CoffeeType,
		ShopName:   req.ShopName,
		Location:   req.Location,
		ImageURL:   req.ImageURL,
		DrinkDate:  drinkDate,
		Mood:       req.Mood,
		Notes:      req.Notes,
		Acidity:    req.Acidity,
		Bitterness: req.Bitterness,
		Sweetness:  req.Sweetness,
		Body:       req.Body,
		Aroma:      req.Aroma,
		Aftertaste: req.Aftertaste,
		AISummary:  aiSummary,
	}

	if err := s.coffeeLogRepo.Create(log, req.FlavorTagIDs); err != nil {
		return nil, err
	}

	created, err := s.coffeeLogRepo.FindByID(log.ID, userID)
	if err != nil {
		return log, nil
	}
	return created, nil
}

func (s *CoffeeLogService) GetByID(id, userID uint) (*model.CoffeeLog, error) {
	log, err := s.coffeeLogRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("coffee log not found")
	}
	return log, nil
}

func (s *CoffeeLogService) GetList(userID uint, pagination *utils.Pagination, month, coffeeType string, tagID uint) ([]model.CoffeeLog, int64, error) {
	return s.coffeeLogRepo.FindList(userID, pagination, month, coffeeType, tagID)
}

func (s *CoffeeLogService) Update(id, userID uint, req UpdateCoffeeLogRequest) (*model.CoffeeLog, error) {
	existing, err := s.coffeeLogRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("coffee log not found")
	}

	acidity := existing.Acidity
	bitterness := existing.Bitterness
	sweetness := existing.Sweetness
	body := existing.Body
	aroma := existing.Aroma
	aftertaste := existing.Aftertaste

	if req.CoffeeName != "" {
		existing.CoffeeName = req.CoffeeName
	}
	if req.CoffeeType != "" {
		existing.CoffeeType = req.CoffeeType
	}
	if req.ShopName != "" {
		existing.ShopName = req.ShopName
	}
	if req.Location != "" {
		existing.Location = req.Location
	}
	if req.ImageURL != "" {
		existing.ImageURL = req.ImageURL
	}
	if req.DrinkDate != "" {
		parsed, err := time.Parse("2006-01-02", req.DrinkDate)
		if err != nil {
			return nil, errors.New("invalid drink_date format")
		}
		existing.DrinkDate = &parsed
	}
	if req.Mood != "" {
		existing.Mood = req.Mood
	}
	if req.Notes != "" {
		existing.Notes = req.Notes
	}
	if req.Acidity != nil {
		acidity = *req.Acidity
		existing.Acidity = *req.Acidity
	}
	if req.Bitterness != nil {
		bitterness = *req.Bitterness
		existing.Bitterness = *req.Bitterness
	}
	if req.Sweetness != nil {
		sweetness = *req.Sweetness
		existing.Sweetness = *req.Sweetness
	}
	if req.Body != nil {
		body = *req.Body
		existing.Body = *req.Body
	}
	if req.Aroma != nil {
		aroma = *req.Aroma
		existing.Aroma = *req.Aroma
	}
	if req.Aftertaste != nil {
		aftertaste = *req.Aftertaste
		existing.Aftertaste = *req.Aftertaste
	}

	if err := s.validateFlavorScores(acidity, bitterness, sweetness, body, aroma, aftertaste); err != nil {
		return nil, err
	}

	needRegenAI := req.Acidity != nil || req.Bitterness != nil || req.Sweetness != nil ||
		req.Body != nil || req.Aroma != nil || req.Aftertaste != nil || len(req.FlavorTagIDs) > 0

	if needRegenAI {
		tagNames := []string{}
		if len(req.FlavorTagIDs) > 0 {
			tags, _ := s.flavorTagRepo.FindByIDs(req.FlavorTagIDs)
			for _, t := range tags {
				tagNames = append(tagNames, t.Name)
			}
		} else {
			for _, t := range existing.FlavorTags {
				tagNames = append(tagNames, t.Name)
			}
		}
		existing.AISummary = s.aiService.GenerateMockSummary(existing.CoffeeType, tagNames, acidity, bitterness, sweetness, body, aroma, aftertaste)
	}

	if err := s.coffeeLogRepo.Update(existing, req.FlavorTagIDs); err != nil {
		return nil, err
	}

	updated, err := s.coffeeLogRepo.FindByID(id, userID)
	if err != nil {
		return existing, nil
	}
	return updated, nil
}

func (s *CoffeeLogService) Delete(id, userID uint) error {
	return s.coffeeLogRepo.Delete(id, userID)
}
