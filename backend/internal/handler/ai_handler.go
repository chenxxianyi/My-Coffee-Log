package handler

import (
	"my-coffee-log/internal/middleware"
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
	if err := service.ValidateFlavorSummaryRequest(req); err != nil {
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

func (h *AIHandler) GetLifestyleQuote(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	resp, err := h.aiService.GenerateLifestyleQuoteForUser(userID.(uint))
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AIHandler) GetFlavorInsight(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	resp, err := h.aiService.GenerateFlavorInsightForUser(userID.(uint))
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}
