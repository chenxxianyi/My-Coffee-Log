package handler

import (
	"strconv"

	"my-coffee-log/internal/middleware"
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"
	"my-coffee-log/internal/utils"

	"github.com/gin-gonic/gin"
)

type CoffeeBeanHandler struct {
	beanService *service.CoffeeBeanService
}

func NewCoffeeBeanHandler(beanService *service.CoffeeBeanService) *CoffeeBeanHandler {
	return &CoffeeBeanHandler{beanService: beanService}
}

func (h *CoffeeBeanHandler) Create(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	var req service.CreateCoffeeBeanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	bean, err := h.beanService.Create(userID, req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, bean)
}

func (h *CoffeeBeanHandler) GetByID(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid bean id")
		return
	}

	bean, err := h.beanService.GetByID(uint(id), userID)
	if err != nil {
		response.Error(c, 40400, err.Error())
		return
	}

	response.Success(c, bean)
}

func (h *CoffeeBeanHandler) GetList(c *gin.Context) {
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

	beans, total, err := h.beanService.GetList(userID, &pagination, search)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.SuccessWithPagination(c, beans, pagination.Page, pagination.PageSize, total)
}

func (h *CoffeeBeanHandler) Update(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid bean id")
		return
	}

	var req service.UpdateCoffeeBeanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	bean, err := h.beanService.Update(uint(id), userID, req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, bean)
}

func (h *CoffeeBeanHandler) Delete(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ErrorBadRequest(c, "invalid bean id")
		return
	}

	if err := h.beanService.Delete(uint(id), userID); err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *CoffeeBeanHandler) GetBeanList(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	beans, err := h.beanService.GetBeanList(userID)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, map[string]interface{}{"beans": beans})
}
