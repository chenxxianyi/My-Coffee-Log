package router

import (
	"time"

	"my-coffee-log/internal/handler"
	"my-coffee-log/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	authHandler      *handler.AuthHandler
	userHandler      *handler.UserHandler
	coffeeLogHandler *handler.CoffeeLogHandler
	statsHandler     *handler.StatsHandler
	aiHandler        *handler.AIHandler
	flavorTagHandler *handler.FlavorTagHandler
	uploadHandler    *handler.UploadHandler
}

func NewRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	coffeeLogHandler *handler.CoffeeLogHandler,
	statsHandler *handler.StatsHandler,
	aiHandler *handler.AIHandler,
	flavorTagHandler *handler.FlavorTagHandler,
	uploadHandler *handler.UploadHandler,
) *Router {
	return &Router{
		authHandler:      authHandler,
		userHandler:      userHandler,
		coffeeLogHandler: coffeeLogHandler,
		statsHandler:     statsHandler,
		aiHandler:        aiHandler,
		flavorTagHandler: flavorTagHandler,
		uploadHandler:    uploadHandler,
	}
}

func (r *Router) Setup(engine *gin.Engine) {
	// Serve uploaded files
	engine.Static("/uploads", "./uploads")

	api := engine.Group("/api/v1")

	// Public routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", r.authHandler.Register)
		auth.POST("/login", r.authHandler.Login)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Users
		users := protected.Group("/users")
		{
			users.GET("/me", r.userHandler.GetCurrentUser)
			users.PUT("/me", r.userHandler.UpdateCurrentUser)
		}

		// Coffee Logs
		coffeeLogs := protected.Group("/coffee-logs")
		{
			coffeeLogs.POST("", r.coffeeLogHandler.Create)
			coffeeLogs.GET("", r.coffeeLogHandler.GetList)
			coffeeLogs.GET("/:id", r.coffeeLogHandler.GetByID)
			coffeeLogs.PUT("/:id", r.coffeeLogHandler.Update)
			coffeeLogs.DELETE("/:id", r.coffeeLogHandler.Delete)
		}

		// Stats
		stats := protected.Group("/stats")
		{
			stats.GET("/overview", r.statsHandler.GetOverview)
			stats.GET("/flavor-profile", r.statsHandler.GetFlavorProfile)
			stats.GET("/monthly", r.statsHandler.GetMonthly)
		}

		// AI
		ai := protected.Group("/ai")
		ai.Use(middleware.MaxBodyBytes(8<<10), middleware.RateLimit(20, time.Minute))
		{
			ai.POST("/flavor-summary", r.aiHandler.GenerateFlavorSummary)
			ai.POST("/lifestyle-quote", r.aiHandler.GetLifestyleQuote)
			ai.POST("/flavor-insight", r.aiHandler.GetFlavorInsight)
		}

		// Uploads
		protected.POST("/upload", r.uploadHandler.UploadFile)

		// Flavor Tags
		flavorTags := protected.Group("/flavor-tags")
		{
			flavorTags.GET("", r.flavorTagHandler.GetAll)
		}
	}
}
