package handler

import (
	"strconv"

	"my-coffee-log/internal/middleware"
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"
	"my-coffee-log/internal/utils"

	"github.com/gin-gonic/gin"
)

type CoffeeLogHandler struct {
	coffeeLogService *service.CoffeeLogService
}

func NewCoffeeLogHandler(coffeeLogService *service.CoffeeLogService) *CoffeeLogHandler {
	return &CoffeeLogHandler{coffeeLogService: coffeeLogService}
}

func (h *CoffeeLogHandler) Create(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	var req service.CreateCoffeeLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	log, err := h.coffeeLogService.Create(userID.(uint), req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, log)
}

func (h *CoffeeLogHandler) GetList(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	pagination := utils.Pagination{
		Page:     1,
		PageSize: 10,
	}

	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			pagination.Page = p
		}
	}
	if pageSizeStr := c.Query("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil {
			pagination.PageSize = ps
		}
	}

	month := c.Query("month")
	coffeeType := c.Query("coffee_type")
	var tagID uint
	if tagIDStr := c.Query("tag_id"); tagIDStr != "" {
		if t, err := strconv.ParseUint(tagIDStr, 10, 32); err == nil {
			tagID = uint(t)
		}
	}

	logs, total, err := h.coffeeLogService.GetList(userID.(uint), &pagination, month, coffeeType, tagID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.SuccessWithPagination(c, logs, pagination.Page, pagination.PageSize, total)
}

func (h *CoffeeLogHandler) GetByID(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid id")
		return
	}

	log, err := h.coffeeLogService.GetByID(uint(id), userID.(uint))
	if err != nil {
		response.ErrorNotFound(c, err.Error())
		return
	}

	response.Success(c, log)
}

func (h *CoffeeLogHandler) Update(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid id")
		return
	}

	var req service.UpdateCoffeeLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	log, err := h.coffeeLogService.Update(uint(id), userID.(uint), req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, log)
}

func (h *CoffeeLogHandler) Delete(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserID)
	if !exists {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid id")
		return
	}

	if err := h.coffeeLogService.Delete(uint(id), userID.(uint)); err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, nil)
}
