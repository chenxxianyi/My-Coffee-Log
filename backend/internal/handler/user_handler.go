package handler

import (
	"my-coffee-log/internal/middleware"
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	user, err := h.userService.GetCurrentUser(userID)
	if err != nil {
		response.ErrorNotFound(c, "user not found")
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) UpdateCurrentUser(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	var req service.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	user, err := h.userService.UpdateUser(userID, req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) CompleteOnboarding(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.ErrorUnauthorized(c, "user not found in context")
		return
	}

	var req service.OnboardingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	user, err := h.userService.CompleteOnboarding(userID, req)
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}

	response.Success(c, user)
}
