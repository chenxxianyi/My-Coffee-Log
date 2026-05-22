package handler

import (
	"my-coffee-log/internal/middleware"
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	statsService *service.StatsService
}

func NewStatsHandler(statsService *service.StatsService) *StatsHandler {
	return &StatsHandler{statsService: statsService}
}

func (h *StatsHandler) GetOverview(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	overview, err := h.statsService.GetOverview(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, overview)
}

func (h *StatsHandler) GetFlavorProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	profile, err := h.statsService.GetFlavorProfile(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, profile)
}

func (h *StatsHandler) GetMonthly(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	monthly, err := h.statsService.GetMonthly(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, monthly)
}

func (h *StatsHandler) GetPersonality(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	personality, err := h.statsService.GetPersonality(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, personality)
}

func (h *StatsHandler) GetMonthlyReview(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	month := c.Query("month")
	review, err := h.statsService.GetMonthlyReview(userID, month)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, review)
}
