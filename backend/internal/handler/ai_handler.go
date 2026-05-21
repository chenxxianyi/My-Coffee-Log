package handler

import (
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"

	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	aiService *service.AIService
}

func NewAIHandler(aiService *service.AIService) *AIHandler {
	return &AIHandler{aiService: aiService}
}

func (h *AIHandler) GenerateFlavorSummary(c *gin.Context) {
	var req service.FlavorSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	resp, err := h.aiService.GenerateFlavorSummary(req)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}
