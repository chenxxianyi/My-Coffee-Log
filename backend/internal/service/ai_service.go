package service

import (
	"fmt"
	"strings"
)

type AIService struct{}

func NewAIService() *AIService {
	return &AIService{}
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
