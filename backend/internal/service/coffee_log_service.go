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
	CoffeeName   string `json:"coffee_name"`
	CoffeeType   string `json:"coffee_type"`
	ShopName     string `json:"shop_name"`
	Location     string `json:"location"`
	ImageURL     string `json:"image_url"`
	DrinkDate    string `json:"drink_date"`
	Mood         string `json:"mood"`
	Notes        string `json:"notes"`
	GenerateAI   bool   `json:"generate_ai"`
	Acidity      int    `json:"acidity"`
	Bitterness   int    `json:"bitterness"`
	Sweetness    int    `json:"sweetness"`
	Body         int    `json:"body"`
	Aroma        int    `json:"aroma"`
	Aftertaste   int    `json:"aftertaste"`
	FlavorTagIDs []uint `json:"flavor_tag_ids"`
}

type UpdateCoffeeLogRequest struct {
	CoffeeName   string  `json:"coffee_name"`
	CoffeeType   string  `json:"coffee_type"`
	ShopName     string  `json:"shop_name"`
	Location     string  `json:"location"`
	ImageURL     string  `json:"image_url"`
	DrinkDate    string  `json:"drink_date"`
	Mood         string  `json:"mood"`
	Notes        string  `json:"notes"`
	GenerateAI   *bool   `json:"generate_ai"`
	Acidity      *int    `json:"acidity"`
	Bitterness   *int    `json:"bitterness"`
	Sweetness    *int    `json:"sweetness"`
	Body         *int    `json:"body"`
	Aroma        *int    `json:"aroma"`
	Aftertaste   *int    `json:"aftertaste"`
	FlavorTagIDs *[]uint `json:"flavor_tag_ids"`
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

func (s *CoffeeLogService) validateFlavorTagIDs(ids []uint) ([]model.FlavorTag, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	tags, err := s.flavorTagRepo.FindByIDs(ids)
	if err != nil || len(tags) != len(ids) {
		return nil, errors.New("one or more flavor tags not found")
	}
	return tags, nil
}

func flavorTagNames(tags []model.FlavorTag) []string {
	names := make([]string, 0, len(tags))
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}

func (s *CoffeeLogService) generateAndStoreAISummary(logID uint, userID uint, req FlavorSummaryRequest) {
	aiResp, err := s.aiService.GenerateFlavorSummary(req)
	if err != nil || aiResp == nil {
		return
	}
	_ = s.coffeeLogRepo.UpdateAISummary(logID, userID, aiResp.Summary)
}

func (s *CoffeeLogService) Create(userID uint, req CreateCoffeeLogRequest) (*model.CoffeeLog, error) {
	if err := s.validateFlavorScores(req.Acidity, req.Bitterness, req.Sweetness, req.Body, req.Aroma, req.Aftertaste); err != nil {
		return nil, err
	}

	tags, err := s.validateFlavorTagIDs(req.FlavorTagIDs)
	if err != nil {
		return nil, err
	}
	tagNames := flavorTagNames(tags)

	var drinkDate *time.Time
	if req.DrinkDate != "" {
		parsed, err := time.Parse("2006-01-02", req.DrinkDate)
		if err != nil {
			return nil, errors.New("invalid drink_date format, expected YYYY-MM-DD")
		}
		drinkDate = &parsed
	}

	aiReq := FlavorSummaryRequest{
		CoffeeName: req.CoffeeName,
		CoffeeType: req.CoffeeType,
		Tags:       tagNames,
		Acidity:    req.Acidity,
		Bitterness: req.Bitterness,
		Sweetness:  req.Sweetness,
		Body:       req.Body,
		Aroma:      req.Aroma,
		Aftertaste: req.Aftertaste,
		Mood:       req.Mood,
		Notes:      req.Notes,
	}
	aiSummary := s.aiService.generateMockSummary(req.CoffeeType, tagNames, req.Acidity, req.Bitterness, req.Sweetness, req.Body, req.Aroma, req.Aftertaste)

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

	if req.GenerateAI && externalAIEnabled() {
		go s.generateAndStoreAISummary(log.ID, userID, aiReq)
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
	aiInputChanged := false

	if req.CoffeeName != "" {
		existing.CoffeeName = req.CoffeeName
		aiInputChanged = true
	}
	if req.CoffeeType != "" {
		existing.CoffeeType = req.CoffeeType
		aiInputChanged = true
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
		aiInputChanged = true
	}
	if req.Notes != "" {
		existing.Notes = req.Notes
		aiInputChanged = true
	}
	if req.Acidity != nil {
		acidity = *req.Acidity
		existing.Acidity = *req.Acidity
		aiInputChanged = true
	}
	if req.Bitterness != nil {
		bitterness = *req.Bitterness
		existing.Bitterness = *req.Bitterness
		aiInputChanged = true
	}
	if req.Sweetness != nil {
		sweetness = *req.Sweetness
		existing.Sweetness = *req.Sweetness
		aiInputChanged = true
	}
	if req.Body != nil {
		body = *req.Body
		existing.Body = *req.Body
		aiInputChanged = true
	}
	if req.Aroma != nil {
		aroma = *req.Aroma
		existing.Aroma = *req.Aroma
		aiInputChanged = true
	}
	if req.Aftertaste != nil {
		aftertaste = *req.Aftertaste
		existing.Aftertaste = *req.Aftertaste
		aiInputChanged = true
	}

	if err := s.validateFlavorScores(acidity, bitterness, sweetness, body, aroma, aftertaste); err != nil {
		return nil, err
	}

	var flavorTagIDs []uint
	tagNames := flavorTagNames(existing.FlavorTags)
	tagIDsProvided := req.FlavorTagIDs != nil
	if tagIDsProvided {
		flavorTagIDs = *req.FlavorTagIDs
		tags, err := s.validateFlavorTagIDs(flavorTagIDs)
		if err != nil {
			return nil, err
		}
		tagNames = flavorTagNames(tags)
	}

	needRegenAI := aiInputChanged || tagIDsProvided
	var aiReq FlavorSummaryRequest
	generateExternalAI := false
	if needRegenAI {
		aiReq = FlavorSummaryRequest{
			CoffeeName: existing.CoffeeName,
			CoffeeType: existing.CoffeeType,
			Tags:       tagNames,
			Acidity:    acidity,
			Bitterness: bitterness,
			Sweetness:  sweetness,
			Body:       body,
			Aroma:      aroma,
			Aftertaste: aftertaste,
			Mood:       existing.Mood,
			Notes:      existing.Notes,
		}
		existing.AISummary = s.aiService.generateMockSummary(existing.CoffeeType, tagNames, acidity, bitterness, sweetness, body, aroma, aftertaste)
		generateExternalAI = req.GenerateAI != nil && *req.GenerateAI && externalAIEnabled()
	}

	if err := s.coffeeLogRepo.Update(existing, flavorTagIDs); err != nil {
		return nil, err
	}

	if generateExternalAI {
		go s.generateAndStoreAISummary(id, userID, aiReq)
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
