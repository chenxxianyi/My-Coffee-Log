package handler

import (
	"my-coffee-log/internal/response"
	"my-coffee-log/internal/repository"

	"github.com/gin-gonic/gin"
)

type FlavorTagHandler struct {
	flavorTagRepo *repository.FlavorTagRepository
}

func NewFlavorTagHandler(flavorTagRepo *repository.FlavorTagRepository) *FlavorTagHandler {
	return &FlavorTagHandler{flavorTagRepo: flavorTagRepo}
}

func (h *FlavorTagHandler) GetAll(c *gin.Context) {
	tags, err := h.flavorTagRepo.FindAll()
	if err != nil {
		response.Error(c, 50000, err.Error())
		return
	}
	response.Success(c, tags)
}
