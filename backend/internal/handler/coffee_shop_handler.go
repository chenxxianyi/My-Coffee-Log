package handler

import (
	"strconv"

	"my-coffee-log/internal/middleware"
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"
	"my-coffee-log/internal/utils"

	"github.com/gin-gonic/gin"
)

type CoffeeShopHandler struct {
	shopService *service.CoffeeShopService
}

func NewCoffeeShopHandler(shopService *service.CoffeeShopService) *CoffeeShopHandler {
	return &CoffeeShopHandler{shopService: shopService}
}

func (h *CoffeeShopHandler) Create(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	var req service.CreateCoffeeShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	shop, err := h.shopService.Create(userID, req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, shop)
}

func (h *CoffeeShopHandler) GetByID(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid shop id")
		return
	}

	shop, err := h.shopService.GetByID(uint(id), userID)
	if err != nil {
		response.Error(c, 40400, err.Error())
		return
	}

	response.Success(c, shop)
}

func (h *CoffeeShopHandler) GetList(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	pagination := utils.Pagination{Page: 1, PageSize: 20}
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			pagination.Page = p
		}
	}
	if pageSizeStr := c.Query("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
			pagination.PageSize = ps
		}
	}

	search := c.Query("search")

	shops, total, err := h.shopService.GetList(userID, &pagination, search)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.SuccessWithPagination(c, shops, pagination.Page, pagination.PageSize, total)
}

func (h *CoffeeShopHandler) Update(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid shop id")
		return
	}

	var req service.UpdateCoffeeShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	shop, err := h.shopService.Update(uint(id), userID, req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, shop)
}

func (h *CoffeeShopHandler) Delete(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid shop id")
		return
	}

	if err := h.shopService.Delete(uint(id), userID); err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *CoffeeShopHandler) GetShopNames(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	names, err := h.shopService.GetShopNames(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, map[string]interface{}{"names": names})
}

func (h *CoffeeShopHandler) GetRelatedLogs(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid shop id")
		return
	}

	pagination := utils.Pagination{Page: 1, PageSize: 20}
	if pageStr := c.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			pagination.Page = p
		}
	}
	if pageSizeStr := c.Query("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
			pagination.PageSize = ps
		}
	}

	logs, total, err := h.shopService.GetRelatedLogs(uint(id), userID, &pagination)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.SuccessWithPagination(c, logs, pagination.Page, pagination.PageSize, total)
}
