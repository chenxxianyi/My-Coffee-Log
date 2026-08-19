package service

import (
	"fmt"
	"sort"
	"time"

	"my-coffee-log/internal/repository"
)

type StatsService struct {
	statsRepo *repository.StatsRepository
}

func NewStatsService(statsRepo *repository.StatsRepository) *StatsService {
	return &StatsService{statsRepo: statsRepo}
}

type FlavorTagItem struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Count int64  `json:"count"`
}

type OverviewResponse struct {
	MonthCount         int64           `json:"month_count"`
	TotalCount         int64           `json:"total_count"`
	FavoriteCoffeeType string          `json:"favorite_coffee_type"`
	FavoriteFlavorTag  string          `json:"favorite_flavor_tag"`
	RecentFlavorTags   []FlavorTagItem `json:"recent_flavor_tags"`
}

// Coffee Personality types
type PersonalityTag struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type PersonalityResponse struct {
	Personalities []PersonalityTag `json:"personalities"`
}

type personalityRule struct {
	slug        string
	title       string
	subtitle    string
	description string
	icon        string
	match       func(ctx personalityContext) bool
	score       func(ctx personalityContext) float64
}

type personalityContext struct {
	favType       string
	flavorProfile *repository.FlavorProfile
	moodTags      []repository.LifestyleTagCount
	sceneTags     []repository.LifestyleTagCount
	pairingTags   []repository.LifestyleTagCount
	totalCount    int64
}

var personalityRules = []personalityRule{
	{
		slug:        "citrus-minimalist",
		title:       "Citrus Minimalist",
		subtitle:    "柑橘极简主义者",
		description: "你偏爱清爽、明亮、酸质突出的咖啡，像一杯晨光中的手冲，轻盈而充满生命力。",
		icon:        "🍋",
		match: func(ctx personalityContext) bool {
			return ctx.flavorProfile != nil && ctx.flavorProfile.Acidity >= 3.5 && ctx.flavorProfile.Body < 3.0
		},
		score: func(ctx personalityContext) float64 {
			if ctx.flavorProfile == nil {
				return 0
			}
			return ctx.flavorProfile.Acidity - ctx.flavorProfile.Body
		},
	},
	{
		slug:        "creamy-comfort-seeker",
		title:       "Creamy Comfort Seeker",
		subtitle:    "奶油舒适追寻者",
		description: "你偏爱奶咖、焦糖、坚果、顺滑口感，每一口都是温暖的拥抱。",
		icon:        "🥛",
		match: func(ctx personalityContext) bool {
			return (ctx.flavorProfile != nil && ctx.flavorProfile.Sweetness >= 3.5 && ctx.flavorProfile.Body >= 3.0) ||
				ctx.favType == "Latte" || ctx.favType == "Cappuccino" || ctx.favType == "Flat White"
		},
		score: func(ctx personalityContext) float64 {
			s := 0.0
			if ctx.flavorProfile != nil {
				s += ctx.flavorProfile.Sweetness + ctx.flavorProfile.Body
			}
			if ctx.favType == "Latte" || ctx.favType == "Cappuccino" || ctx.favType == "Flat White" {
				s += 3
			}
			return s
		},
	},
	{
		slug:        "slow-morning-brewer",
		title:       "Slow Morning Brewer",
		subtitle:    "慢晨手冲师",
		description: "常在早晨记录手冲，偏好安静、轻盈的风味，咖啡是你开启慢生活的仪式。",
		icon:        "🌅",
		match: func(ctx personalityContext) bool {
			hasMorning := false
			for _, t := range ctx.sceneTags {
				if t.Tag == "Morning" {
					hasMorning = true
					break
				}
			}
			return hasMorning && (ctx.favType == "Pour Over" || ctx.favType == "Cold Brew")
		},
		score: func(ctx personalityContext) float64 {
			s := 0.0
			if ctx.favType == "Pour Over" {
				s += 2
			}
			for _, t := range ctx.sceneTags {
				if t.Tag == "Morning" {
					s += float64(t.Count)
				}
			}
			return s
		},
	},
	{
		slug:        "urban-latte-lover",
		title:       "Urban Latte Lover",
		subtitle:    "都市拿铁爱好者",
		description: "偏好咖啡店、拿铁、通勤场景，城市中的每一杯拿铁都是你的生活节拍。",
		icon:        "🏙️",
		match: func(ctx personalityContext) bool {
			hasCafe := false
			hasOffice := false
			for _, t := range ctx.sceneTags {
				if t.Tag == "Cafe" {
					hasCafe = true
				}
				if t.Tag == "Office" {
					hasOffice = true
				}
			}
			return (hasCafe || hasOffice) && (ctx.favType == "Latte" || ctx.favType == "Americano")
		},
		score: func(ctx personalityContext) float64 {
			s := 0.0
			if ctx.favType == "Latte" || ctx.favType == "Americano" {
				s += 2
			}
			for _, t := range ctx.sceneTags {
				if t.Tag == "Cafe" || t.Tag == "Office" {
					s += float64(t.Count)
				}
			}
			return s
		},
	},
	{
		slug:        "bold-espresso-purist",
		title:       "Bold Espresso Purist",
		subtitle:    "浓郁浓缩纯粹主义者",
		description: "你追求浓烈、醇厚、苦甜交织的深度体验，浓缩是你对咖啡最真诚的致敬。",
		icon:        "☕",
		match: func(ctx personalityContext) bool {
			return (ctx.flavorProfile != nil && ctx.flavorProfile.Bitterness >= 3.5 && ctx.flavorProfile.Body >= 3.5) ||
				ctx.favType == "Espresso" || ctx.favType == "Dirty"
		},
		score: func(ctx personalityContext) float64 {
			s := 0.0
			if ctx.flavorProfile != nil {
				s += ctx.flavorProfile.Bitterness + ctx.flavorProfile.Body
			}
			if ctx.favType == "Espresso" || ctx.favType == "Dirty" {
				s += 3
			}
			return s
		},
	},
	{
		slug:        "rainy-day-reader",
		title:       "Rainy Day Reader",
		subtitle:    "雨日阅读者",
		description: "阴雨天、独处、阅读——你的咖啡时光总是与安静为伴，一杯咖啡一本书就是整个世界。",
		icon:        "📖",
		match: func(ctx personalityContext) bool {
			hasRainy := false
			hasBook := false
			hasAlone := false
			for _, t := range ctx.moodTags {
				if t.Tag == "Rainy" {
					hasRainy = true
				}
			}
			for _, t := range ctx.pairingTags {
				if t.Tag == "Book" {
					hasBook = true
				}
				if t.Tag == "Alone" {
					hasAlone = true
				}
			}
			return hasRainy || (hasBook && hasAlone)
		},
		score: func(ctx personalityContext) float64 {
			s := 0.0
			for _, t := range ctx.moodTags {
				if t.Tag == "Rainy" || t.Tag == "Slow" {
					s += float64(t.Count)
				}
			}
			for _, t := range ctx.pairingTags {
				if t.Tag == "Book" || t.Tag == "Alone" {
					s += float64(t.Count)
				}
			}
			return s
		},
	},
	{
		slug:        "social-weekend-explorer",
		title:       "Social Weekend Explorer",
		subtitle:    "社交周末探索者",
		description: "周末、朋友、甜点——你的咖啡是社交的媒介，每一杯都连接着一段美好时光。",
		icon:        "🎉",
		match: func(ctx personalityContext) bool {
			hasWeekend := false
			hasFriends := false
			for _, t := range ctx.sceneTags {
				if t.Tag == "Weekend" {
					hasWeekend = true
				}
			}
			for _, t := range ctx.pairingTags {
				if t.Tag == "Friends" {
					hasFriends = true
				}
			}
			return hasWeekend && hasFriends
		},
		score: func(ctx personalityContext) float64 {
			s := 0.0
			for _, t := range ctx.sceneTags {
				if t.Tag == "Weekend" || t.Tag == "Travel" {
					s += float64(t.Count)
				}
			}
			for _, t := range ctx.pairingTags {
				if t.Tag == "Friends" || t.Tag == "Dessert" {
					s += float64(t.Count)
				}
			}
			return s
		},
	},
	{
		slug:        "productive-hustler",
		title:       "Productive Hustler",
		subtitle:    "高效奋斗者",
		description: "专注、高效、工作——咖啡是你的燃料，每一杯都推动你向目标更进一步。",
		icon:        "🚀",
		match: func(ctx personalityContext) bool {
			hasProductive := false
			hasWork := false
			for _, t := range ctx.moodTags {
				if t.Tag == "Productive" || t.Tag == "Focused" {
					hasProductive = true
				}
			}
			for _, t := range ctx.pairingTags {
				if t.Tag == "Work" {
					hasWork = true
				}
			}
			return hasProductive || hasWork
		},
		score: func(ctx personalityContext) float64 {
			s := 0.0
			for _, t := range ctx.moodTags {
				if t.Tag == "Productive" || t.Tag == "Focused" {
					s += float64(t.Count)
				}
			}
			for _, t := range ctx.pairingTags {
				if t.Tag == "Work" {
					s += float64(t.Count)
				}
			}
			return s
		},
	},
}

func (s *StatsService) GetPersonality(userID uint) (*PersonalityResponse, error) {
	favType, _ := s.statsRepo.GetFavoriteCoffeeType(userID)
	profile, _ := s.statsRepo.GetFlavorProfile(userID)
	totalCount, _ := s.statsRepo.GetTotalCount(userID)
	moodTags, _ := s.statsRepo.GetLifestyleTagCounts(userID, "mood_tags")
	sceneTags, _ := s.statsRepo.GetLifestyleTagCounts(userID, "scene_tags")
	pairingTags, _ := s.statsRepo.GetLifestyleTagCounts(userID, "pairing_tags")

	ctx := personalityContext{
		favType:       favType,
		flavorProfile: profile,
		moodTags:      moodTags,
		sceneTags:     sceneTags,
		pairingTags:   pairingTags,
		totalCount:    totalCount,
	}

	type scored struct {
		tag   PersonalityTag
		score float64
	}
	var matched []scored
	for _, rule := range personalityRules {
		if rule.match(ctx) {
			matched = append(matched, scored{
				tag: PersonalityTag{
					Slug:        rule.slug,
					Title:       rule.title,
					Subtitle:    rule.subtitle,
					Description: rule.description,
					Icon:        rule.icon,
				},
				score: rule.score(ctx),
			})
		}
	}

	sort.Slice(matched, func(i, j int) bool { return matched[i].score > matched[j].score })

	if len(matched) > 3 {
		matched = matched[:3]
	}

	tags := make([]PersonalityTag, 0, len(matched))
	for _, m := range matched {
		tags = append(tags, m.tag)
	}

	// Fallback: if no personality matched but user has logs, assign based on fav type
	if len(tags) == 0 && totalCount > 0 {
		fallback := PersonalityTag{
			Slug:        "coffee-explorer",
			Title:       "Coffee Explorer",
			Subtitle:    "咖啡探索者",
			Description: "你正在探索属于自己的咖啡世界，每一杯都是一次新的发现。",
			Icon:        "🧭",
		}
		tags = append(tags, fallback)
	}

	return &PersonalityResponse{Personalities: tags}, nil
}

// StatsMeta holds sample count and threshold info for stats responses
type StatsMeta struct {
	SampleCount      int64 `json:"sample_count"`
	ValidSensoryCount int64 `json:"valid_sensory_count"`
	DateFrom         string `json:"date_from,omitempty"`
	DateTo           string `json:"date_to,omitempty"`
	Threshold        int   `json:"threshold"`
	IsReady          bool  `json:"is_ready"`
}

// StatsResponse wraps data with meta information
type StatsResponse struct {
	Data interface{} `json:"data"`
	Meta StatsMeta   `json:"meta"`
}

// Threshold constants for different insights
const (
	ThresholdBasicTypePreference = 3
	ThresholdFlavorRadar         = 3
	ThresholdWeeklyTrend         = 5
	ThresholdCoffeePersonality   = 8
	ThresholdMonthlyComparison   = 6
)

func (s *StatsService) GetOverview(userID uint) (*OverviewResponse, error) {
	monthCount, err := s.statsRepo.GetMonthCount(userID)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.statsRepo.GetTotalCount(userID)
	if err != nil {
		return nil, err
	}

	favType, err := s.statsRepo.GetFavoriteCoffeeType(userID)
	if err != nil {
		return nil, err
	}

	favTag, err := s.statsRepo.GetFavoriteFlavorTag(userID)
	if err != nil {
		return nil, err
	}

	recentTags, err := s.statsRepo.GetRecentFlavorTags(userID, 5)
	if err != nil {
		return nil, err
	}

	tagItems := make([]FlavorTagItem, 0, len(recentTags))
	for _, t := range recentTags {
		tagItems = append(tagItems, FlavorTagItem{
			Name:  t.Name,
			Label: t.Label,
			Count: t.Count,
		})
	}

	return &OverviewResponse{
		MonthCount:         monthCount,
		TotalCount:         totalCount,
		FavoriteCoffeeType: favType,
		FavoriteFlavorTag:  favTag,
		RecentFlavorTags:   tagItems,
	}, nil
}

// GetOverviewWithMeta returns overview with sample count and threshold meta
func (s *StatsService) GetOverviewWithMeta(userID uint) (*StatsResponse, error) {
	overview, err := s.GetOverview(userID)
	if err != nil {
		return nil, err
	}

	validSensory, _ := s.statsRepo.GetValidSensoryCount(userID)
	dateFrom, dateTo := s.getDateRange(userID)

	return &StatsResponse{
		Data: overview,
		Meta: StatsMeta{
			SampleCount:       overview.TotalCount,
			ValidSensoryCount: validSensory,
			DateFrom:          dateFrom,
			DateTo:            dateTo,
			Threshold:         ThresholdBasicTypePreference,
			IsReady:           overview.TotalCount >= int64(ThresholdBasicTypePreference),
		},
	}, nil
}

func (s *StatsService) GetFlavorProfile(userID uint) (*repository.FlavorProfile, error) {
	return s.statsRepo.GetFlavorProfile(userID)
}

func (s *StatsService) GetMonthly(userID uint) ([]repository.MonthlyCount, error) {
	return s.statsRepo.GetMonthlyCounts(userID)
}

func (s *StatsService) GetRecentMood(userID uint) (string, error) {
	return s.statsRepo.GetRecentMood(userID)
}

// ---- Monthly Review ----

type MonthlyReviewFlavorTag struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Count int64  `json:"count"`
}

type MonthlyReviewCoffeeType struct {
	CoffeeType string `json:"coffee_type"`
	Count      int64  `json:"count"`
}

type MonthlyReviewCoffeeName struct {
	CoffeeName string `json:"coffee_name"`
	Count      int64  `json:"count"`
}

type MonthlyReviewLifestyleTag struct {
	Tag   string `json:"tag"`
	Count int64  `json:"count"`
}

type MonthlyReviewFlavorProfile struct {
	Acidity    float64 `json:"acidity"`
	Bitterness float64 `json:"bitterness"`
	Sweetness  float64 `json:"sweetness"`
	Body       float64 `json:"body"`
	Aroma      float64 `json:"aroma"`
	Aftertaste float64 `json:"aftertaste"`
}

type MonthlyReviewResponse struct {
	Month         string                      `json:"month"`
	Count         int64                       `json:"count"`
	FavCoffeeType string                      `json:"favorite_coffee_type"`
	CoffeeTypes   []MonthlyReviewCoffeeType   `json:"coffee_types"`
	FlavorTags    []MonthlyReviewFlavorTag    `json:"flavor_tags"`
	CoffeeNames   []MonthlyReviewCoffeeName   `json:"coffee_names"`
	TopWeekday    *int                        `json:"top_weekday"`
	MoodTags      []MonthlyReviewLifestyleTag `json:"mood_tags"`
	SceneTags     []MonthlyReviewLifestyleTag `json:"scene_tags"`
	PairingTags   []MonthlyReviewLifestyleTag `json:"pairing_tags"`
	FlavorProfile *MonthlyReviewFlavorProfile `json:"flavor_profile"`
	Keywords      []string                    `json:"keywords"`
	AISummary     string                      `json:"ai_summary"`
}

var weekdayNames = map[int]string{
	1: "周日", 2: "周一", 3: "周二", 4: "周三",
	5: "周四", 6: "周五", 7: "周六",
}

func (s *StatsService) GetMonthlyReview(userID uint, month string) (*MonthlyReviewResponse, error) {
	if month == "" {
		now := time.Now()
		month = fmt.Sprintf("%04d-%02d", now.Year(), now.Month())
	}

	count, err := s.statsRepo.GetMonthCountByMonth(userID, month)
	if err != nil {
		return nil, err
	}

	favType, _ := s.statsRepo.GetFavoriteCoffeeTypeByMonth(userID, month)

	typeCounts, _ := s.statsRepo.GetTopCoffeeTypesByMonth(userID, month, 5)
	coffeeTypes := make([]MonthlyReviewCoffeeType, 0, len(typeCounts))
	for _, tc := range typeCounts {
		coffeeTypes = append(coffeeTypes, MonthlyReviewCoffeeType{
			CoffeeType: tc.CoffeeType,
			Count:      tc.Count,
		})
	}

	tagCounts, _ := s.statsRepo.GetTopFlavorTagsByMonth(userID, month, 6)
	flavorTags := make([]MonthlyReviewFlavorTag, 0, len(tagCounts))
	for _, tc := range tagCounts {
		flavorTags = append(flavorTags, MonthlyReviewFlavorTag{
			Name:  tc.Name,
			Label: tc.Label,
			Count: tc.Count,
		})
	}

	nameCounts, _ := s.statsRepo.GetTopCoffeeNamesByMonth(userID, month, 5)
	coffeeNames := make([]MonthlyReviewCoffeeName, 0, len(nameCounts))
	for _, nc := range nameCounts {
		coffeeNames = append(coffeeNames, MonthlyReviewCoffeeName{
			CoffeeName: nc.CoffeeName,
			Count:      nc.Count,
		})
	}

	var topWeekday *int
	weekdayResult, _ := s.statsRepo.GetTopDrinkWeekdayByMonth(userID, month)
	if weekdayResult != nil {
		topWeekday = &weekdayResult.Weekday
	}

	moodTagCounts, _ := s.statsRepo.GetLifestyleTagCountsByMonth(userID, month, "mood_tags")
	moodTags := make([]MonthlyReviewLifestyleTag, 0, len(moodTagCounts))
	for _, tc := range moodTagCounts {
		moodTags = append(moodTags, MonthlyReviewLifestyleTag{Tag: tc.Tag, Count: tc.Count})
	}

	sceneTagCounts, _ := s.statsRepo.GetLifestyleTagCountsByMonth(userID, month, "scene_tags")
	sceneTags := make([]MonthlyReviewLifestyleTag, 0, len(sceneTagCounts))
	for _, tc := range sceneTagCounts {
		sceneTags = append(sceneTags, MonthlyReviewLifestyleTag{Tag: tc.Tag, Count: tc.Count})
	}

	pairingTagCounts, _ := s.statsRepo.GetLifestyleTagCountsByMonth(userID, month, "pairing_tags")
	pairingTags := make([]MonthlyReviewLifestyleTag, 0, len(pairingTagCounts))
	for _, tc := range pairingTagCounts {
		pairingTags = append(pairingTags, MonthlyReviewLifestyleTag{Tag: tc.Tag, Count: tc.Count})
	}

	profile, _ := s.statsRepo.GetFlavorProfileByMonth(userID, month)
	var flavorProfile *MonthlyReviewFlavorProfile
	if profile != nil {
		flavorProfile = &MonthlyReviewFlavorProfile{
			Acidity:    profile.Acidity,
			Bitterness: profile.Bitterness,
			Sweetness:  profile.Sweetness,
			Body:       profile.Body,
			Aroma:      profile.Aroma,
			Aftertaste: profile.Aftertaste,
		}
	}

	// Generate keywords from top data
	keywords := s.generateMonthlyKeywords(favType, flavorTags, coffeeNames, moodTags, sceneTags, topWeekday)

	return &MonthlyReviewResponse{
		Month:         month,
		Count:         count,
		FavCoffeeType: favType,
		CoffeeTypes:   coffeeTypes,
		FlavorTags:    flavorTags,
		CoffeeNames:   coffeeNames,
		TopWeekday:    topWeekday,
		MoodTags:      moodTags,
		SceneTags:     sceneTags,
		PairingTags:   pairingTags,
		FlavorProfile: flavorProfile,
		Keywords:      keywords,
		AISummary:     "",
	}, nil
}

func (s *StatsService) generateMonthlyKeywords(favType string, flavorTags []MonthlyReviewFlavorTag, coffeeNames []MonthlyReviewCoffeeName, moodTags []MonthlyReviewLifestyleTag, sceneTags []MonthlyReviewLifestyleTag, topWeekday *int) []string {
	typesCN := map[string]string{
		"Pour Over": "手冲", "Latte": "拿铁", "Americano": "美式",
		"Cold Brew": "冷萃", "Espresso": "浓缩", "Dirty": "脏咖啡",
		"Cappuccino": "卡布奇诺", "Flat White": "馥芮白",
	}

	tagsCN := map[string]string{
		"floral": "花香", "citrus": "柑橘", "berry": "莓果", "nutty": "坚果",
		"chocolate": "巧克力", "caramel": "焦糖", "creamy": "奶油", "winey": "酒香",
		"smoky": "烟熏", "herbal": "草本",
	}

	moodCN := map[string]string{
		"Calm": "平静", "Energetic": "愉悦", "Reflective": "沉浸", "Tired": "疲惫",
		"Focused": "专注", "Happy": "开心", "Rainy": "阴雨", "Slow": "慢活", "Productive": "高效",
	}

	sceneCN := map[string]string{
		"Morning": "早晨", "Office": "办公", "Weekend": "周末", "Cafe": "咖啡馆",
		"Travel": "旅行", "Home": "居家", "Study": "学习",
	}

	var keywords []string

	if favType != "" {
		if cn, ok := typesCN[favType]; ok {
			keywords = append(keywords, cn+"爱好者")
		}
	}

	if len(flavorTags) > 0 {
		if cn, ok := tagsCN[flavorTags[0].Name]; ok {
			keywords = append(keywords, cn+"风味")
		}
	}

	if len(coffeeNames) > 0 {
		keywords = append(keywords, coffeeNames[0].CoffeeName)
	}

	if len(moodTags) > 0 {
		if cn, ok := moodCN[moodTags[0].Tag]; ok {
			keywords = append(keywords, cn+"时光")
		}
	}

	if len(sceneTags) > 0 {
		if cn, ok := sceneCN[sceneTags[0].Tag]; ok {
			keywords = append(keywords, cn+"场景")
		}
	}

	if topWeekday != nil {
		if name, ok := weekdayNames[*topWeekday]; ok {
			keywords = append(keywords, name+"常客")
		}
	}

	if len(keywords) > 6 {
		keywords = keywords[:6]
	}

	return keywords
}

// getDateRange returns the earliest and latest drink dates for a user
func (s *StatsService) getDateRange(userID uint) (string, string) {
	return s.statsRepo.GetDateRange(userID)
}

// ---- Weekly Review ----

type WeeklyReviewCoffeeType struct {
	CoffeeType string `json:"coffee_type"`
	Count      int64  `json:"count"`
}

type WeeklyReviewFlavorTag struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Count int64  `json:"count"`
}

type WeeklyReviewLifestyleTag struct {
	Tag   string `json:"tag"`
	Count int64  `json:"count"`
}

type WeeklyReviewFlavorProfile struct {
	Acidity    float64 `json:"acidity"`
	Bitterness float64 `json:"bitterness"`
	Sweetness  float64 `json:"sweetness"`
	Body       float64 `json:"body"`
	Aroma      float64 `json:"aroma"`
	Aftertaste float64 `json:"aftertaste"`
}

type WeeklyReviewResponse struct {
	Week          string                        `json:"week"`
	WeekStart     string                        `json:"week_start"`
	WeekEnd       string                        `json:"week_end"`
	Count         int64                         `json:"count"`
	FavCoffeeType string                        `json:"favorite_coffee_type"`
	CoffeeTypes   []WeeklyReviewCoffeeType      `json:"coffee_types"`
	FlavorTags    []WeeklyReviewFlavorTag       `json:"flavor_tags"`
	MoodTags      []WeeklyReviewLifestyleTag    `json:"mood_tags"`
	SceneTags     []WeeklyReviewLifestyleTag    `json:"scene_tags"`
	PairingTags   []WeeklyReviewLifestyleTag    `json:"pairing_tags"`
	FlavorProfile *WeeklyReviewFlavorProfile    `json:"flavor_profile"`
	Trend         string                        `json:"trend"`
	Memory        string                        `json:"memory"`
}

func (s *StatsService) GetWeeklyReview(userID uint, week string) (*WeeklyReviewResponse, error) {
	if week == "" {
		now := time.Now()
		year, weekNum := now.ISOWeek()
		week = fmt.Sprintf("%d-W%02d", year, weekNum)
	}

	// Parse week string to get date range
	var year, weekNum int
	_, err := fmt.Sscanf(week, "%d-W%d", &year, &weekNum)
	if err != nil {
		return nil, fmt.Errorf("invalid week format, expected YYYY-Wxx")
	}

	// Calculate week start (Monday) and end (Sunday)
	weekStart := getWeekStart(year, weekNum)
	weekEnd := weekStart.AddDate(0, 0, 6)

	startStr := weekStart.Format("2006-01-02")
	endStr := weekEnd.Format("2006-01-02")

	count, err := s.statsRepo.GetCountByDateRange(userID, startStr, endStr)
	if err != nil {
		return nil, err
	}

	favType, _ := s.statsRepo.GetFavoriteCoffeeTypeByDateRange(userID, startStr, endStr)

	typeCounts, _ := s.statsRepo.GetTopCoffeeTypesByDateRange(userID, startStr, endStr, 5)
	coffeeTypes := make([]WeeklyReviewCoffeeType, 0, len(typeCounts))
	for _, tc := range typeCounts {
		coffeeTypes = append(coffeeTypes, WeeklyReviewCoffeeType{
			CoffeeType: tc.CoffeeType,
			Count:      tc.Count,
		})
	}

	tagCounts, _ := s.statsRepo.GetTopFlavorTagsByDateRange(userID, startStr, endStr, 6)
	flavorTags := make([]WeeklyReviewFlavorTag, 0, len(tagCounts))
	for _, tc := range tagCounts {
		flavorTags = append(flavorTags, WeeklyReviewFlavorTag{
			Name:  tc.Name,
			Label: tc.Label,
			Count: tc.Count,
		})
	}

	moodTagCounts, _ := s.statsRepo.GetLifestyleTagCountsByDateRange(userID, startStr, endStr, "mood_tags")
	moodTags := make([]WeeklyReviewLifestyleTag, 0, len(moodTagCounts))
	for _, tc := range moodTagCounts {
		moodTags = append(moodTags, WeeklyReviewLifestyleTag{Tag: tc.Tag, Count: tc.Count})
	}

	sceneTagCounts, _ := s.statsRepo.GetLifestyleTagCountsByDateRange(userID, startStr, endStr, "scene_tags")
	sceneTags := make([]WeeklyReviewLifestyleTag, 0, len(sceneTagCounts))
	for _, tc := range sceneTagCounts {
		sceneTags = append(sceneTags, WeeklyReviewLifestyleTag{Tag: tc.Tag, Count: tc.Count})
	}

	pairingTagCounts, _ := s.statsRepo.GetLifestyleTagCountsByDateRange(userID, startStr, endStr, "pairing_tags")
	pairingTags := make([]WeeklyReviewLifestyleTag, 0, len(pairingTagCounts))
	for _, tc := range pairingTagCounts {
		pairingTags = append(pairingTags, WeeklyReviewLifestyleTag{Tag: tc.Tag, Count: tc.Count})
	}

	profile, _ := s.statsRepo.GetFlavorProfileByDateRange(userID, startStr, endStr)
	var flavorProfile *WeeklyReviewFlavorProfile
	if profile != nil {
		flavorProfile = &WeeklyReviewFlavorProfile{
			Acidity:    profile.Acidity,
			Bitterness: profile.Bitterness,
			Sweetness:  profile.Sweetness,
			Body:       profile.Body,
			Aroma:      profile.Aroma,
			Aftertaste: profile.Aftertaste,
		}
	}

	// Generate trend text
	trend := s.generateWeeklyTrend(count, favType, moodTags, sceneTags)

	// Generate memory text
	memory := s.generateWeeklyMemory(count, coffeeTypes, flavorTags)

	return &WeeklyReviewResponse{
		Week:          week,
		WeekStart:     startStr,
		WeekEnd:       endStr,
		Count:         count,
		FavCoffeeType: favType,
		CoffeeTypes:   coffeeTypes,
		FlavorTags:    flavorTags,
		MoodTags:      moodTags,
		SceneTags:     sceneTags,
		PairingTags:   pairingTags,
		FlavorProfile: flavorProfile,
		Trend:         trend,
		Memory:        memory,
	}, nil
}

func getWeekStart(year, week int) time.Time {
	// ISO week: week 1 is the week with the first Thursday of the year
	// January 4th is always in week 1
	jan4 := time.Date(year, 1, 4, 0, 0, 0, 0, time.UTC)
	startOfJan4Week := jan4.AddDate(0, 0, -int(jan4.Weekday()-time.Monday))
	if startOfJan4Week.Weekday() == time.Sunday {
		startOfJan4Week = startOfJan4Week.AddDate(0, 0, -6)
	}
	weekStart := startOfJan4Week.AddDate(0, 0, (week-1)*7)
	return weekStart
}

func (s *StatsService) generateWeeklyTrend(count int64, favType string, moodTags []WeeklyReviewLifestyleTag, sceneTags []WeeklyReviewLifestyleTag) string {
	if count == 0 {
		return "本周还没有记录，快来记录你的第一杯吧！"
	}

	typesCN := map[string]string{
		"Pour Over": "手冲", "Latte": "拿铁", "Americano": "美式",
		"Cold Brew": "冷萃", "Espresso": "浓缩", "Dirty": "脏咖啡",
	}

	moodCN := map[string]string{
		"Calm": "平静", "Focused": "专注", "Tired": "疲惫",
		"Happy": "开心", "Rainy": "阴雨", "Slow": "慢活", "Productive": "高效",
	}

	sceneCN := map[string]string{
		"Morning": "早晨", "Office": "办公", "Weekend": "周末",
		"Cafe": "咖啡馆", "Travel": "旅行", "Home": "居家", "Study": "学习",
	}

	trend := fmt.Sprintf("本周记录了 %d 杯咖啡", count)

	if favType != "" {
		if cn, ok := typesCN[favType]; ok {
			trend += fmt.Sprintf("，最常喝%s", cn)
		}
	}

	if len(moodTags) > 0 {
		if cn, ok := moodCN[moodTags[0].Tag]; ok {
			trend += fmt.Sprintf("，心情多为%s", cn)
		}
	}

	if len(sceneTags) > 0 {
		if cn, ok := sceneCN[sceneTags[0].Tag]; ok {
			trend += fmt.Sprintf("，常在%s场景", cn)
		}
	}

	return trend + "。"
}

func (s *StatsService) generateWeeklyMemory(count int64, coffeeTypes []WeeklyReviewCoffeeType, flavorTags []WeeklyReviewFlavorTag) string {
	if count == 0 {
		return ""
	}

	if len(coffeeTypes) > 0 && len(flavorTags) > 0 {
		typesCN := map[string]string{
			"Pour Over": "手冲", "Latte": "拿铁", "Americano": "美式",
			"Cold Brew": "冷萃", "Espresso": "浓缩", "Dirty": "脏咖啡",
		}
		tagsCN := map[string]string{
			"floral": "花香", "citrus": "柑橘", "berry": "莓果", "nutty": "坚果",
			"chocolate": "巧克力", "caramel": "焦糖", "creamy": "奶油", "winey": "酒香",
			"smoky": "烟熏", "herbal": "草本",
		}

		typeCN := coffeeTypes[0].CoffeeType
		tagCN := flavorTags[0].Name
		if cn, ok := typesCN[typeCN]; ok {
			typeCN = cn
		}
		if cn, ok := tagsCN[tagCN]; ok {
			tagCN = cn
		}

		return fmt.Sprintf("本周最爱%s + %s风味", typeCN, tagCN)
	}

	return fmt.Sprintf("本周共记录了 %d 杯咖啡", count)
}

// ---- Record Progress ----

type RecordProgressResponse struct {
	RecordID         uint   `json:"record_id"`
	MonthCupCount    int64  `json:"month_cup_count"`
	TotalCount       int64  `json:"total_count"`
	IsFirstRecord    bool   `json:"is_first_record"`
	NextInsightName  string `json:"next_insight_name"`
	NextInsightDelta int    `json:"next_insight_delta"`
}

func (s *StatsService) GetRecordProgress(userID uint, recordID uint) (*RecordProgressResponse, error) {
	monthCount, err := s.statsRepo.GetMonthCount(userID)
	if err != nil {
		return nil, err
	}
	totalCount, err := s.statsRepo.GetTotalCount(userID)
	if err != nil {
		return nil, err
	}

	isFirst := totalCount <= 1

	// Determine next insight and how many more records needed
	nextInsightName := ""
	nextDelta := 0

	if totalCount < int64(ThresholdBasicTypePreference) {
		nextInsightName = "基础类型偏好"
		nextDelta = ThresholdBasicTypePreference - int(totalCount)
	} else if totalCount < int64(ThresholdFlavorRadar) {
		nextInsightName = "初步风味雷达"
		nextDelta = ThresholdFlavorRadar - int(totalCount)
	} else if totalCount < int64(ThresholdWeeklyTrend) {
		nextInsightName = "周度趋势"
		nextDelta = ThresholdWeeklyTrend - int(totalCount)
	} else if totalCount < int64(ThresholdCoffeePersonality) {
		nextInsightName = "咖啡人格"
		nextDelta = ThresholdCoffeePersonality - int(totalCount)
	} else {
		nextInsightName = ""
		nextDelta = 0
	}

	return &RecordProgressResponse{
		RecordID:         recordID,
		MonthCupCount:    monthCount,
		TotalCount:       totalCount,
		IsFirstRecord:    isFirst,
		NextInsightName:  nextInsightName,
		NextInsightDelta: nextDelta,
	}, nil
}
