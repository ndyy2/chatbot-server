// main.go
package main

import (
	"ai-assistant/config"
	"ai-assistant/controllers"
	"ai-assistant/models"
	"ai-assistant/repositories"
	"ai-assistant/routes"
	"ai-assistant/services"
	"ai-assistant/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize databases
	mariadb, err := config.InitMariaDB(cfg.MariaDBURI)
	if err != nil {
		log.Fatal("MariaDB connection failed:", err)
	}
	
	mongodb, err := config.InitMongoDB(cfg.MongoDBURI)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	// Auto-migrate MariaDB models
	mariadb.AutoMigrate(
		&models.SystemSetting{},
		&models.User{},
		&models.AuditLog{},
	)

	// Initialize repositories
	systemRepo := repositories.NewSystemRepository(mariadb)
	memoryRepo := repositories.NewMemoryRepository(mongodb.Database("ai_assistant"))

	// Initialize services
	groqClient := utils.NewGroqClient(cfg.GroqAPIKey, "meta-llama/llama-4-scout-17b-16e-instruct")
	aiService := services.NewAIService(groqClient, systemRepo, memoryRepo)

	// Initialize controllers
	aiController := controllers.NewAIController(aiService)

	// Setup Echo
	e := echo.New()

	// Register routes
	routes.RegisterAPIRoutes(e, aiController)

	// Start server
	log.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}