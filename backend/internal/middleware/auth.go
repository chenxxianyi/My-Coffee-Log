package middleware

import (
	"strings"

	"my-coffee-log/internal/response"
	"my-coffee-log/internal/utils"

	"github.com/gin-gonic/gin"
)

const ContextUserID = "user_id"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorUnauthorized(c, "Authorization header is required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.ErrorUnauthorized(c, "Authorization header format must be Bearer <token>")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.ErrorUnauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set(ContextUserID, claims.UserID)
		c.Next()
	}
}

// GetUserID safely extracts uint userID from gin context.
// Returns (0, false) if not found or wrong type.
func GetUserID(c *gin.Context) (uint, bool) {
	val, exists := c.Get(ContextUserID)
	if !exists {
		return 0, false
	}
	uid, ok := val.(uint)
	return uid, ok
}
