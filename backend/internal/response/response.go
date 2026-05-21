package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationData struct {
	List       interface{} `json:"list"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func SuccessWithPagination(c *gin.Context, list interface{}, page, pageSize int, total int64) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data: PaginationData{
			List: list,
			Pagination: Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    total,
			},
		},
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func ErrorBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    40001,
		Message: message,
		Data:    nil,
	})
}

func ErrorUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    40100,
		Message: message,
		Data:    nil,
	})
}

func ErrorForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, Response{
		Code:    40300,
		Message: message,
		Data:    nil,
	})
}

func ErrorNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Code:    40400,
		Message: message,
		Data:    nil,
	})
}

func ErrorInternal(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    50000,
		Message: message,
		Data:    nil,
	})
}
