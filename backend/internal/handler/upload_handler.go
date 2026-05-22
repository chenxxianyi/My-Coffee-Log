package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"my-coffee-log/internal/response"

	"github.com/gin-gonic/gin"
)

var allowedMIMETypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
	"image/gif":  true,
}

var allowedExtensions = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true, ".gif": true,
}

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
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

	if file.Size > 5*1024*1024 {
		response.ErrorBadRequest(c, "File size exceeds limit (5MB)")
		return
	}

	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		response.ErrorBadRequest(c, "Invalid file format. Allowed formats: jpg, jpeg, png, webp, gif")
		return
	}

	// Validate MIME type from header
	contentType := file.Header.Get("Content-Type")
	if !allowedMIMETypes[contentType] {
		response.ErrorBadRequest(c, "Invalid MIME type")
		return
	}

	// Validate actual file content via magic bytes
	src, err := file.Open()
	if err != nil {
		response.ErrorInternal(c, "Failed to read file")
		return
	}
	buf := make([]byte, 512)
	n, _ := src.Read(buf)
	src.Close()
	if n > 0 {
		detected := http.DetectContentType(buf[:n])
		if !allowedMIMETypes[detected] {
			response.ErrorBadRequest(c, "File content does not match allowed image type")
			return
		}
	}

	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		_ = os.MkdirAll("./uploads", 0755)
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	targetPath := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, targetPath); err != nil {
		response.ErrorInternal(c, "Failed to save file: "+err.Error())
		return
	}

	fileURL := "/uploads/" + filename
	response.Success(c, gin.H{
		"url": fileURL,
	})
}
