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
