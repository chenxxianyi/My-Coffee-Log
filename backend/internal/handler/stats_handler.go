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
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	overview, err := h.statsService.GetOverview(userID.(uint))
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, overview)
}

func (h *StatsHandler) GetFlavorProfile(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	profile, err := h.statsService.GetFlavorProfile(userID.(uint))
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, profile)
}

func (h *StatsHandler) GetMonthly(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	monthly, err := h.statsService.GetMonthly(userID.(uint))
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, monthly)
}
