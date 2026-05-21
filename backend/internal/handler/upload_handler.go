package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"my-coffee-log/internal/response"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	// Create uploads directory if not exists
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		_ = os.MkdirAll("./uploads", 0755)
	}
	return &UploadHandler{}
}

func (h *UploadHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorBadRequest(c, "No file uploaded")
		return
	}

	// Validate size (max 5MB)
	if file.Size > 5*1024*1024 {
		response.ErrorBadRequest(c, "File size exceeds limit (5MB)")
		return
	}

	// Validate extension
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" && ext != ".gif" {
		response.ErrorBadRequest(c, "Invalid file format. Allowed formats: jpg, jpeg, png, webp, gif")
		return
	}

	// Ensure uploads directory exists again securely
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		_ = os.MkdirAll("./uploads", 0755)
	}

	// Generate clean unique filename
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	targetPath := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, targetPath); err != nil {
		response.ErrorInternal(c, "Failed to save file: "+err.Error())
		return
	}

	// Return absolute path mapping under dev domain/relative path
	fileURL := "/uploads/" + filename
	response.Success(c, gin.H{
		"url": fileURL,
	})
}
