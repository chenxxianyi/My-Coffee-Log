package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"my-coffee-log/internal/response"

	"github.com/gin-gonic/gin"
)

type rateLimitBucket struct {
	count       atomic.Int64
	windowStart atomic.Int64 // unix nanoseconds
}

var rateLimitStore sync.Map

func init() {
	go cleanupRateLimitStore()
}

func cleanupRateLimitStore() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		now := time.Now().UnixNano()
		rateLimitStore.Range(func(key, value interface{}) bool {
			b := value.(*rateLimitBucket)
			windowNs := b.windowStart.Load()
			// Remove entries older than 2 minutes (conservative upper bound for any window)
			if now-windowNs > int64(2*time.Minute) {
				rateLimitStore.Delete(key)
			}
			return true
		})
	}
}

func MaxBodyBytes(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		c.Next()
	}
}

func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	windowNs := int64(window)
	return func(c *gin.Context) {
		key := c.ClientIP()
		if userID, exists := c.Get(ContextUserID); exists {
			key = ContextUserID + ":" + toString(userID)
		}

		now := time.Now().UnixNano()
		value, _ := rateLimitStore.LoadOrStore(key, &rateLimitBucket{})
		bucket := value.(*rateLimitBucket)

		start := bucket.windowStart.Load()
		if now-start >= windowNs {
			// Window expired, reset
			bucket.windowStart.CompareAndSwap(start, now)
			bucket.count.Store(1)
			c.Next()
			return
		}

		count := bucket.count.Add(1)
		if count > int64(limit) {
			c.JSON(http.StatusTooManyRequests, response.Response{Code: 42900, Message: "rate limit exceeded", Data: nil})
			c.Abort()
			return
		}

		c.Next()
	}
}

func toString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
