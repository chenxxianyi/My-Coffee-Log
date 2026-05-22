package main

import (
	"log"

	"my-coffee-log/internal/config"
	"my-coffee-log/internal/database"
	"my-coffee-log/internal/handler"
	"my-coffee-log/internal/middleware"
	"my-coffee-log/internal/model"
	"my-coffee-log/internal/repository"
	"my-coffee-log/internal/router"
	"my-coffee-log/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Load config
	config.LoadConfig()

	// Init logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer logger.Sync()

	// Init MySQL
	database.InitMySQL()

	// Init Redis
	database.InitRedis()

	// Auto migrate
	if err := database.DB.AutoMigrate(
		&model.User{},
		&model.CoffeeLog{},
		&model.FlavorTag{},
		&model.CardTemplate{},
	); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// Seed default flavor tags
	flavorTagRepo := repository.NewFlavorTagRepository(database.DB)
	seedFlavorTags(flavorTagRepo)

	// Init repositories
	userRepo := repository.NewUserRepository(database.DB)
	coffeeLogRepo := repository.NewCoffeeLogRepository(database.DB)
	statsRepo := repository.NewStatsRepository(database.DB)

	// Init services
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	statsService := service.NewStatsService(statsRepo)
	aiService := service.NewAIService(statsService)
	coffeeLogService := service.NewCoffeeLogService(coffeeLogRepo, flavorTagRepo, aiService)

	// Init handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	coffeeLogHandler := handler.NewCoffeeLogHandler(coffeeLogService)
	statsHandler := handler.NewStatsHandler(statsService)
	aiHandler := handler.NewAIHandler(aiService)
	flavorTagHandler := handler.NewFlavorTagHandler(flavorTagRepo)
	uploadHandler := handler.NewUploadHandler()

	// Init router
	r := router.NewRouter(authHandler, userHandler, coffeeLogHandler, statsHandler, aiHandler, flavorTagHandler, uploadHandler)

	// Setup Gin
	engine := gin.New()
	engine.Use(middleware.RecoveryMiddleware())
	engine.Use(middleware.CORSMiddleware())
	engine.Use(middleware.LoggerMiddleware(logger))

	r.Setup(engine)

	// Start server
	port := config.AppConfig.AppPort
	if port == "" {
		port = "8080"
	}
	logger.Info("Server starting", zap.String("port", port))
	if err := engine.Run(":" + port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func seedFlavorTags(repo *repository.FlavorTagRepository) {
	defaultTags := []model.FlavorTag{
		{Name: "floral", Label: "����", Color: "#D4A5A5"},
		{Name: "citrus", Label: "����", Color: "#F5C156"},
		{Name: "berry", Label: "ݮ��", Color: "#B5443E"},
		{Name: "nutty", Label: "���", Color: "#A67B5B"},
		{Name: "chocolate", Label: "�ɿ���", Color: "#5C3317"},
		{Name: "caramel", Label: "����", Color: "#D4A017"},
		{Name: "creamy", Label: "����", Color: "#F5E6CC"},
		{Name: "winey", Label: "����", Color: "#722F37"},
		{Name: "smoky", Label: "��Ѭ", Color: "#4A4A4A"},
		{Name: "herbal", Label: "�ݱ�", Color: "#6B8E23"},
	}
	if err := repo.Seed(defaultTags); err != nil {
		log.Printf("Warning: failed to seed flavor tags: %v", err)
	} else {
		log.Println("Flavor tags seeded successfully")
	}
}
