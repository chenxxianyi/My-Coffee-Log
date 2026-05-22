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
	shopHandler      *handler.CoffeeShopHandler
	beanHandler      *handler.CoffeeBeanHandler
}

func NewRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	coffeeLogHandler *handler.CoffeeLogHandler,
	statsHandler *handler.StatsHandler,
	aiHandler *handler.AIHandler,
	flavorTagHandler *handler.FlavorTagHandler,
	uploadHandler *handler.UploadHandler,
	shopHandler *handler.CoffeeShopHandler,
	beanHandler *handler.CoffeeBeanHandler,
) *Router {
	return &Router{
		authHandler:      authHandler,
		userHandler:      userHandler,
		coffeeLogHandler: coffeeLogHandler,
		statsHandler:     statsHandler,
		aiHandler:        aiHandler,
		flavorTagHandler: flavorTagHandler,
		uploadHandler:    uploadHandler,
		shopHandler:      shopHandler,
		beanHandler:      beanHandler,
	}
}

func (r *Router) Setup(engine *gin.Engine) {
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
			stats.GET("/personality", r.statsHandler.GetPersonality)
			stats.GET("/monthly-review", r.statsHandler.GetMonthlyReview)
		}

		// AI
		ai := protected.Group("/ai")
		ai.Use(middleware.MaxBodyBytes(8<<10), middleware.RateLimit(20, time.Minute))
		{
			ai.POST("/flavor-summary", r.aiHandler.GenerateFlavorSummary)
			ai.POST("/lifestyle-quote", r.aiHandler.GetLifestyleQuote)
			ai.POST("/flavor-insight", r.aiHandler.GetFlavorInsight)
			ai.GET("/monthly-review", r.aiHandler.GetMonthlyReview)
			ai.GET("/status", r.aiHandler.GetAIStatus)
			ai.POST("/share-copy", r.aiHandler.GenerateShareCopy)
			ai.POST("/coffee-profile", r.aiHandler.GenerateCoffeeProfile)
			ai.POST("/preference-insight", r.aiHandler.GeneratePreferenceInsight)
		}

		// Uploads (rate limited)
		protected.POST("/upload", middleware.RateLimit(10, time.Minute), r.uploadHandler.UploadFile)

		// Serve uploaded files behind auth
		protected.Static("/uploads", "./uploads")

		// Flavor Tags
		flavorTags := protected.Group("/flavor-tags")
		{
			flavorTags.GET("", r.flavorTagHandler.GetAll)
		}

		// Coffee Shops
		shops := protected.Group("/coffee-shops")
		{
			shops.POST("", r.shopHandler.Create)
			shops.GET("", r.shopHandler.GetList)
			shops.GET("/names", r.shopHandler.GetShopNames)
			shops.GET("/:id", r.shopHandler.GetByID)
			shops.PUT("/:id", r.shopHandler.Update)
			shops.DELETE("/:id", r.shopHandler.Delete)
			shops.GET("/:id/logs", r.shopHandler.GetRelatedLogs)
		}

		// Coffee Beans
		beans := protected.Group("/coffee-beans")
		{
			beans.POST("", r.beanHandler.Create)
			beans.GET("", r.beanHandler.GetList)
			beans.GET("/list", r.beanHandler.GetBeanList)
			beans.GET("/:id", r.beanHandler.GetByID)
			beans.PUT("/:id", r.beanHandler.Update)
			beans.DELETE("/:id", r.beanHandler.Delete)
		}
	}
}
