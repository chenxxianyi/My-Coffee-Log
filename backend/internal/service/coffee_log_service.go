package service

import (
	"encoding/json"
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
	shopService   *CoffeeShopService
	beanService   *CoffeeBeanService
	userService   *UserService
}

func NewCoffeeLogService(coffeeLogRepo *repository.CoffeeLogRepository, flavorTagRepo *repository.FlavorTagRepository, aiService *AIService, shopService *CoffeeShopService, beanService *CoffeeBeanService, userService *UserService) *CoffeeLogService {
	return &CoffeeLogService{
		coffeeLogRepo: coffeeLogRepo,
		flavorTagRepo: flavorTagRepo,
		aiService:     aiService,
		shopService:   shopService,
		beanService:   beanService,
		userService:   userService,
	}
}

type CreateCoffeeLogRequest struct {
	CoffeeName      string   `json:"coffee_name" binding:"max=120"`
	CoffeeType      string   `json:"coffee_type" binding:"required,max=50"`
	ShopName        string   `json:"shop_name" binding:"max=255"`
	Location        string   `json:"location" binding:"max=255"`
	ImageURL        string   `json:"image_url" binding:"max=500"`
	DrinkDate       string   `json:"drink_date" binding:"max=20"`
	Mood            string   `json:"mood" binding:"max=50"`
	Notes           string   `json:"notes" binding:"max=2000"`
	GenerateAI      bool     `json:"generate_ai"`
	Acidity         int      `json:"acidity"`
	Bitterness      int      `json:"bitterness"`
	Sweetness       int      `json:"sweetness"`
	Body            int      `json:"body"`
	Aroma           int      `json:"aroma"`
	Aftertaste      int      `json:"aftertaste"`
	FlavorTagIDs    []uint   `json:"flavor_tag_ids"`
	MoodTags        []string `json:"mood_tags"`
	SceneTags       []string `json:"scene_tags"`
	PairingTags     []string `json:"pairing_tags"`
	BeanID          *uint    `json:"bean_id"`
	BeanName        string   `json:"bean_name" binding:"max=120"`
	BrewRatio       string   `json:"brew_ratio" binding:"max=50"`
	WaterTemp       string   `json:"water_temp" binding:"max=50"`
	GrindSize       string   `json:"grind_size" binding:"max=50"`

	// Data quality fields (v2)
	RecordMode       string `json:"record_mode" binding:"omitempty,oneof=quick detailed"`     // quick | detailed
	CoffeeNameSource string `json:"coffee_name_source" binding:"omitempty,max=30"`              // user_input | system_suggested | empty
	NotesSource      string `json:"notes_source" binding:"omitempty,max=30"`                   // user_input | ai_generated | empty
	ShopSource       string `json:"shop_source" binding:"omitempty,max=30"`                    // user_input | recent_reuse | empty
	SensoryRecorded  bool   `json:"sensory_recorded"`
	SourceLogID      *uint  `json:"source_log_id"`
	IsTestData       bool   `json:"is_test_data"`
}

type UpdateCoffeeLogRequest struct {
	CoffeeName      string    `json:"coffee_name"`
	CoffeeType      string    `json:"coffee_type"`
	ShopName        string    `json:"shop_name"`
	Location        string    `json:"location"`
	ImageURL        string    `json:"image_url"`
	DrinkDate       string    `json:"drink_date"`
	Mood            string    `json:"mood"`
	Notes           string    `json:"notes"`
	GenerateAI      *bool     `json:"generate_ai"`
	Acidity         *int      `json:"acidity"`
	Bitterness      *int      `json:"bitterness"`
	Sweetness       *int      `json:"sweetness"`
	Body            *int      `json:"body"`
	Aroma           *int      `json:"aroma"`
	Aftertaste      *int      `json:"aftertaste"`
	FlavorTagIDs    *[]uint   `json:"flavor_tag_ids"`
	MoodTags        *[]string `json:"mood_tags"`
	SceneTags       *[]string `json:"scene_tags"`
	PairingTags     *[]string `json:"pairing_tags"`
	BeanID          *uint     `json:"bean_id"`
	BrewRatio       *string   `json:"brew_ratio"`
	WaterTemp       *string   `json:"water_temp"`
	GrindSize       *string   `json:"grind_size"`

	// Allow clearing optional fields using a presence flag
	ClearShopName        bool `json:"clear_shop_name"`
	ClearNotes           bool `json:"clear_notes"`
	ClearMoodTags        bool `json:"clear_mood_tags"`
	ClearSceneTags       bool `json:"clear_scene_tags"`
	ClearPairingTags     bool `json:"clear_pairing_tags"`
	ClearFlavorTags      bool `json:"clear_flavor_tags"`
	ClearLocation        bool `json:"clear_location"`
	ClearBrewRatio       bool `json:"clear_brew_ratio"`
	ClearWaterTemp       bool `json:"clear_water_temp"`
	ClearGrindSize       bool `json:"clear_grind_size"`
	ClearBeanID          bool `json:"clear_bean_id"`
	ClearSensoryRecorded bool `json:"clear_sensory_recorded"`

	// Data quality updates
	SensoryRecorded *bool `json:"sensory_recorded_update"`
}

const maxLifestyleTags = 5
const maxLifestyleTagLen = 30

var validMoodTags = map[string]bool{
	"Calm": true, "Focused": true, "Tired": true, "Happy": true,
	"Rainy": true, "Slow": true, "Productive": true,
}
var validSceneTags = map[string]bool{
	"Morning": true, "Office": true, "Weekend": true, "Cafe": true,
	"Travel": true, "Home": true, "Study": true,
}
var validPairingTags = map[string]bool{
	"Book": true, "Music": true, "Work": true, "Dessert": true,
	"Alone": true, "Friends": true,
}

func validateLifestyleTags(tags []string, validSet map[string]bool, maxCount int, label string) error {
	if len(tags) > maxCount {
		return fmt.Errorf("too many %s (max %d)", label, maxCount)
	}
	for _, t := range tags {
		if len(t) > maxLifestyleTagLen {
			return fmt.Errorf("%s tag too long", label)
		}
		if !validSet[t] {
			return fmt.Errorf("invalid %s tag: %s", label, t)
		}
	}
	return nil
}

func tagsToJSON(tags []string) string {
	if len(tags) == 0 {
		return ""
	}
	b, err := json.Marshal(tags)
	if err != nil {
		return "[]"
	}
	return string(b)
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
	// Validate sensory scores only when sensory_recorded is true
	if req.SensoryRecorded {
		if err := s.validateFlavorScores(req.Acidity, req.Bitterness, req.Sweetness, req.Body, req.Aroma, req.Aftertaste); err != nil {
			return nil, err
		}
	}

	if err := validateLifestyleTags(req.MoodTags, validMoodTags, maxLifestyleTags, "mood"); err != nil {
		return nil, err
	}
	if err := validateLifestyleTags(req.SceneTags, validSceneTags, maxLifestyleTags, "scene"); err != nil {
		return nil, err
	}
	if err := validateLifestyleTags(req.PairingTags, validPairingTags, maxLifestyleTags, "pairing"); err != nil {
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

	// Determine record mode
	recordMode := req.RecordMode
	if recordMode == "" {
		recordMode = "quick"
	}

	// Auto-generate coffee name if empty (system_suggested)
	coffeeName := req.CoffeeName
	coffeeNameSource := req.CoffeeNameSource
	if coffeeName == "" {
		coffeeName = generateAutoCoffeeName(req.CoffeeType)
		if coffeeNameSource == "" {
			coffeeNameSource = "system_suggested"
		}
	}
	if coffeeNameSource == "" {
		coffeeNameSource = "user_input"
	}

	// Determine shop source
	shopSource := req.ShopSource
	if shopSource == "" {
		if req.ShopName == "" {
			shopSource = "empty"
		} else {
			shopSource = "user_input"
		}
	}

	// Determine notes source
	notesSource := req.NotesSource
	if notesSource == "" {
		if req.Notes == "" {
			notesSource = "empty"
		} else {
			notesSource = "user_input"
		}
	}

	aiReq := FlavorSummaryRequest{
		CoffeeName: coffeeName,
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

	// Only generate mock AI summary if user explicitly asked for it
	var aiSummary string
	if req.GenerateAI {
		aiSummary = s.aiService.generateMockSummary(req.CoffeeType, tagNames, req.Acidity, req.Bitterness, req.Sweetness, req.Body, req.Aroma, req.Aftertaste)
	}

	// Resolve bean_id: either from direct reference or inline bean_name
	var beanID *uint
	if s.beanService != nil {
		resolvedID, err := s.beanService.EnsureBeanForLog(userID, req.BeanID, req.BeanName)
		if err == nil && resolvedID != nil {
			beanID = resolvedID
		}
	}

	log := &model.CoffeeLog{
		UserID:          userID,
		CoffeeName:      coffeeName,
		CoffeeType:      req.CoffeeType,
		ShopName:        req.ShopName,
		Location:        req.Location,
		ImageURL:        req.ImageURL,
		DrinkDate:       drinkDate,
		Mood:            req.Mood,
		Notes:           req.Notes,
		Acidity:         req.Acidity,
		Bitterness:      req.Bitterness,
		Sweetness:       req.Sweetness,
		Body:            req.Body,
		Aroma:           req.Aroma,
		Aftertaste:      req.Aftertaste,
		AISummary:       aiSummary,
		MoodTags:        tagsToJSON(req.MoodTags),
		SceneTags:       tagsToJSON(req.SceneTags),
		PairingTags:     tagsToJSON(req.PairingTags),
		BeanID:          beanID,
		BrewRatio:       req.BrewRatio,
		WaterTemp:       req.WaterTemp,
		GrindSize:       req.GrindSize,
		RecordMode:      recordMode,
		CoffeeNameSource: coffeeNameSource,
		NotesSource:      notesSource,
		ShopSource:       shopSource,
		SensoryRecorded:  req.SensoryRecorded,
		SourceLogID:      req.SourceLogID,
		IsTestData:       req.IsTestData,
	}

	if err := s.coffeeLogRepo.Create(log, req.FlavorTagIDs); err != nil {
		return nil, err
	}

	if req.GenerateAI && externalAIEnabled() {
		go s.generateAndStoreAISummary(log.ID, userID, aiReq)
	}

	// Auto-create or update coffee shop record only when user provided a shop name
	if s.shopService != nil && req.ShopName != "" {
		_ = s.shopService.EnsureShopForLog(userID, req.ShopName)
	}

	// Mark first record time for onboarding
	if s.userService != nil {
		_ = s.userService.MarkFirstRecord(userID)
	}

	created, err := s.coffeeLogRepo.FindByID(log.ID, userID)
	if err != nil {
		return log, nil
	}
	return created, nil
}

// generateAutoCoffeeName generates a coffee name from type when user doesn't provide one
func generateAutoCoffeeName(coffeeType string) string {
	typeMap := map[string]string{
		"Pour Over":  "手冲咖啡",
		"Latte":      "拿铁",
		"Americano":  "美式咖啡",
		"Cold Brew":  "冷萃咖啡",
		"Espresso":   "浓缩咖啡",
		"Dirty":      "Dirty 咖啡",
		"Cappuccino": "卡布奇诺",
		"Flat White": "馥芮白",
	}
	if name, ok := typeMap[coffeeType]; ok {
		return name
	}
	return "咖啡"
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
		existing.CoffeeNameSource = "user_input"
		aiInputChanged = true
	}
	if req.CoffeeType != "" {
		existing.CoffeeType = req.CoffeeType
		aiInputChanged = true
	}

	// Handle shop_name: update or clear
	if req.ClearShopName {
		existing.ShopName = ""
		existing.ShopSource = "empty"
	} else if req.ShopName != "" {
		existing.ShopName = req.ShopName
		existing.ShopSource = "user_input"
	}

	if req.ClearLocation {
		existing.Location = ""
	} else if req.Location != "" {
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

	// Handle notes: update or clear
	if req.ClearNotes {
		existing.Notes = ""
		existing.NotesSource = "empty"
	} else if req.Notes != "" {
		existing.Notes = req.Notes
		existing.NotesSource = "user_input"
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
	if req.MoodTags != nil {
		if err := validateLifestyleTags(*req.MoodTags, validMoodTags, maxLifestyleTags, "mood"); err != nil {
			return nil, err
		}
		existing.MoodTags = tagsToJSON(*req.MoodTags)
		aiInputChanged = true
	}
	if req.ClearMoodTags {
		existing.MoodTags = ""
		aiInputChanged = true
	}
	if req.SceneTags != nil {
		if err := validateLifestyleTags(*req.SceneTags, validSceneTags, maxLifestyleTags, "scene"); err != nil {
			return nil, err
		}
		existing.SceneTags = tagsToJSON(*req.SceneTags)
		aiInputChanged = true
	}
	if req.ClearSceneTags {
		existing.SceneTags = ""
		aiInputChanged = true
	}
	if req.PairingTags != nil {
		if err := validateLifestyleTags(*req.PairingTags, validPairingTags, maxLifestyleTags, "pairing"); err != nil {
			return nil, err
		}
		existing.PairingTags = tagsToJSON(*req.PairingTags)
		aiInputChanged = true
	}
	if req.ClearPairingTags {
		existing.PairingTags = ""
		aiInputChanged = true
	}

	// Handle sensory_recorded updates
	if req.SensoryRecorded != nil {
		existing.SensoryRecorded = *req.SensoryRecorded
	} else if req.ClearSensoryRecorded {
		existing.SensoryRecorded = false
	}

	if existing.SensoryRecorded {
		if err := s.validateFlavorScores(acidity, bitterness, sweetness, body, aroma, aftertaste); err != nil {
			return nil, err
		}
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
	if req.ClearFlavorTags {
		flavorTagIDs = []uint{}
		tagNames = []string{}
		aiInputChanged = true
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
		// Only regenerate mock AI if user explicitly requested
		if req.GenerateAI != nil && *req.GenerateAI {
			existing.AISummary = s.aiService.generateMockSummary(existing.CoffeeType, tagNames, acidity, bitterness, sweetness, body, aroma, aftertaste)
			generateExternalAI = externalAIEnabled()
		}
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
