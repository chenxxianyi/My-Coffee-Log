package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"my-coffee-log/internal/response"

	"github.com/gin-gonic/gin"
)

type rateLimitBucket struct {
	Count       int
	WindowStart time.Time
}

var rateLimitStore sync.Map
var rateLimitMu sync.Mutex

func MaxBodyBytes(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		c.Next()
	}
}

func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		if userID, exists := c.Get(ContextUserID); exists {
			key = ContextUserID + ":" + toString(userID)
		}

		now := time.Now()
		rateLimitMu.Lock()
		value, _ := rateLimitStore.LoadOrStore(key, &rateLimitBucket{Count: 0, WindowStart: now})
		bucket := value.(*rateLimitBucket)

		if now.Sub(bucket.WindowStart) >= window {
			bucket.Count = 0
			bucket.WindowStart = now
		}

		bucket.Count++
		if bucket.Count > limit {
			rateLimitMu.Unlock()
			c.JSON(http.StatusTooManyRequests, response.Response{Code: 42900, Message: "rate limit exceeded", Data: nil})
			c.Abort()
			return
		}
		rateLimitMu.Unlock()

		c.Next()
	}
}

func toString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
