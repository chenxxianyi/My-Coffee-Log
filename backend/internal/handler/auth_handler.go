package handler

import (
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	resp, err := h.authService.Register(req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequest(c, err.Error())
		return
	}

	resp, err := h.authService.Login(req)
	if err != nil {
		response.Error(c, 40001, err.Error())
		return
	}

	response.Success(c, resp)
}
