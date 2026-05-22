package handler

import (
	"fmt"
	"my-coffee-log/internal/config"
	"my-coffee-log/internal/middleware"
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"
	"time"

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
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	resp, err := h.aiService.GenerateLifestyleQuoteForUser(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AIHandler) GetFlavorInsight(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	resp, err := h.aiService.GenerateFlavorInsightForUser(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AIHandler) GetMonthlyReview(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	month := c.Query("month")
	if month == "" {
		now := time.Now()
		month = fmt.Sprintf("%04d-%02d", now.Year(), now.Month())
	}

	resp, err := h.aiService.GenerateMonthlyReviewForUser(userID, month)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AIHandler) GetAIStatus(c *gin.Context) {
	enabled := service.ExternalAIEnabled()
	model := ""
	if enabled && config.AppConfig != nil {
		model = config.AppConfig.OpenAIModel
		if model == "" {
			model = "deepseek-chat"
		}
	}
	response.Success(c, gin.H{
		"enabled": enabled,
		"model":   model,
	})
}

func (h *AIHandler) GenerateShareCopy(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	var req service.ShareCopyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	resp, err := h.aiService.GenerateShareCopy(req, userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AIHandler) GenerateCoffeeProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	resp, err := h.aiService.GenerateCoffeeProfileForUser(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AIHandler) GeneratePreferenceInsight(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	resp, err := h.aiService.GeneratePreferenceInsightForUser(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, resp)
}
