package service

import (
	"errors"
	"fmt"
	"strings"
)

type AIService struct {
	statsService *StatsService
}

func NewAIService(statsService *StatsService) *AIService {
	return &AIService{statsService: statsService}
}

type FlavorSummaryRequest struct {
	CoffeeType string   `json:"coffee_type" binding:"required"`
	Tags       []string `json:"tags"`
	Acidity    int      `json:"acidity"`
	Bitterness int      `json:"bitterness"`
	Sweetness  int      `json:"sweetness"`
	Body       int      `json:"body"`
	Aroma      int      `json:"aroma"`
	Aftertaste int      `json:"aftertaste"`
}

type FlavorSummaryResponse struct {
	Summary string `json:"summary"`
}

func (s *AIService) GenerateMockSummary(coffeeType string, tags []string, acidity, bitterness, sweetness, body, aroma, aftertaste int) string {
	typesCN := map[string]string{
		"Pour Over": "手冲手作",
		"Latte":     "醇香拿铁",
		"Americano": "经典美式",
		"Cold Brew": "冰润冷萃",
		"Espresso":  "意式浓缩",
		"Dirty":     "Dirty",
	}
	brewName, ok := typesCN[coffeeType]
	if !ok {
		brewName = "手作咖啡"
	}

	var acidProse string
	if acidity >= 4 {
		acidProse = "明亮高扬的酸质如剔透的水晶，"
	} else if acidity <= 1 {
		acidProse = "其酸度轻抚味蕾而无任何波澜，温润醇静，"
	} else {
		acidProse = "酸质圆润温和，平衡感极佳，"
	}

	var bodyProse string
	if body >= 4 {
		bodyProse = "厚重饱满的身体质感在口中如丝绒滑过，"
	} else {
		bodyProse = "如清冽薄雾般的轻盈体态，呼吸般纯净，"
	}

	var sweetProse string
	if sweetness >= 4 {
		sweetProse = "焦糖和红糖般的馥郁甜美深深融化，"
	} else {
		sweetProse = "带有极细微且令人愉悦的淡甜尾韵，"
	}

	var aromaProse string
	if aroma >= 4 {
		aromaProse = "其高亢昂扬的花果香气在空气中悠长铺展。"
	} else {
		aromaProse = "其气味幽雅恬静，低调温存。"
	}

	tagLabel := ""
	if len(tags) > 0 {
		tagLabel = "伴随着" + strings.Join(tags, "与") + "的风味印记，"
	}

	return fmt.Sprintf("这杯由你精心记录的咖啡，在%s中，%s%s%s%s%s这一切风味印记，正温柔应和着你此时此刻的心境，愿这杯温热的咖啡，为你带来惬意和安宁的一刻。",
		brewName, acidProse, bodyProse, sweetProse, tagLabel, aromaProse)
}

func (s *AIService) GenerateFlavorSummary(req FlavorSummaryRequest) (*FlavorSummaryResponse, error) {
	summary := s.GenerateMockSummary(req.CoffeeType, req.Tags, req.Acidity, req.Bitterness, req.Sweetness, req.Body, req.Aroma, req.Aftertaste)
	return &FlavorSummaryResponse{Summary: summary}, nil
}

type LifestyleQuoteRequest struct {
	MonthCount int      `json:"month_count"`
	FavType    string   `json:"favorite_coffee_type"`
	FlavorTags []string `json:"flavor_tags"`
	RecentMood string   `json:"recent_mood"`
}

type LifestyleQuoteResponse struct {
	Quote string `json:"quote"`
}

func (s *AIService) GenerateLifestyleQuote(req LifestyleQuoteRequest) (*LifestyleQuoteResponse, error) {
	if req.MonthCount < 0 {
		req.MonthCount = 0
	}
	if req.MonthCount > 500 {
		req.MonthCount = 500
	}
	if len(req.FlavorTags) > 10 {
		req.FlavorTags = req.FlavorTags[:10]
	}

	typesCN := map[string]string{
		"Pour Over":  "手冲",
		"Latte":      "拿铁",
		"Americano":  "美式",
		"Cold Brew":  "冷萃",
		"Espresso":   "浓缩",
		"Dirty":      "脏咖啡",
		"Cappuccino": "卡布奇诺",
		"Flat White": "馥芮白",
	}
	brewName, ok := typesCN[req.FavType]
	if !ok {
		brewName = "咖啡"
	}

	tagsCN := map[string]string{
		"floral":    "花香",
		"citrus":    "柑橘",
		"berry":     "莓果",
		"nutty":     "坚果",
		"chocolate": "巧克力",
		"caramel":   "焦糖",
		"creamy":    "奶油",
		"winey":     "酒香",
		"smoky":     "烟熏",
		"herbal":    "草本",
	}

	moodCN := map[string]string{
		"Calm":       "平静",
		"Energetic":  "愉悦",
		"Reflective": "沉浸",
		"Tired":      "疲惫",
	}
	moodLabel, ok := moodCN[req.RecentMood]
	if !ok {
		moodLabel = "宁静"
	}

	var tagLabels []string
	for _, t := range req.FlavorTags {
		if label, ok := tagsCN[t]; ok {
			tagLabels = append(tagLabels, label)
		}
	}

	if req.MonthCount == 0 {
		return &LifestyleQuoteResponse{
			Quote: "咖啡的香气是属于清晨与安静午后的赞美诗。期待你写下今日的第一篇味觉手账，每一杯都值得被温柔记录。",
		}, nil
	}

	var countPhrase string
	if req.MonthCount >= 20 {
		countPhrase = fmt.Sprintf("这个月你已经与咖啡相伴了%d次，每一杯都是生活节奏中安静的注脚。", req.MonthCount)
	} else if req.MonthCount >= 10 {
		countPhrase = fmt.Sprintf("本月%d杯咖啡，是你与生活之间温柔的对话。", req.MonthCount)
	} else if req.MonthCount >= 3 {
		countPhrase = fmt.Sprintf("这个月你记录了%d杯咖啡，每一杯都是属于自己的安静时刻。", req.MonthCount)
	} else {
		countPhrase = fmt.Sprintf("你已经记录了本月第%d杯，这是属于你自己的生活手账。", req.MonthCount)
	}

	var flavorPhrase string
	if len(tagLabels) > 0 {
		flavorPhrase = fmt.Sprintf("你偏爱的%s风味，像一条贯穿味觉记忆的线索，", strings.Join(tagLabels, "与"))
	} else {
		flavorPhrase = "你每一次记录的风味印记，"
	}

	quote := fmt.Sprintf("%s%s与%s的心境彼此呼应，在%s的日常里，每一杯咖啡都是生活赠予的温柔时刻。",
		countPhrase, flavorPhrase, moodLabel, brewName)

	return &LifestyleQuoteResponse{Quote: quote}, nil
}

func (s *AIService) GenerateLifestyleQuoteForUser(userID uint) (*LifestyleQuoteResponse, error) {
	if s.statsService == nil {
		return nil, errors.New("stats service is not configured")
	}

	overview, err := s.statsService.GetOverview(userID)
	if err != nil {
		return nil, err
	}

	recentMood, err := s.statsService.GetRecentMood(userID)
	if err != nil {
		return nil, err
	}
	if recentMood == "" {
		recentMood = "Calm"
	}

	tags := make([]string, 0, len(overview.RecentFlavorTags))
	for _, tag := range overview.RecentFlavorTags {
		tags = append(tags, tag.Name)
	}

	return s.GenerateLifestyleQuote(LifestyleQuoteRequest{
		MonthCount: int(overview.MonthCount),
		FavType:    overview.FavoriteCoffeeType,
		FlavorTags: tags,
		RecentMood: recentMood,
	})
}
