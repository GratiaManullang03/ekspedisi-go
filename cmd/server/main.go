package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/GratiaManullang03/ekspedisi-go/internal/config"
	"github.com/GratiaManullang03/ekspedisi-go/internal/domain/repository"
	"github.com/GratiaManullang03/ekspedisi-go/internal/domain/usecase"
	"github.com/GratiaManullang03/ekspedisi-go/internal/handler"
	"github.com/GratiaManullang03/ekspedisi-go/internal/middleware"
	"github.com/GratiaManullang03/ekspedisi-go/pkg/database"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	db := database.InitDB(cfg)

	// Setup repositories
	shippingRepo := repository.NewShippingRepository(db)

	// Setup use cases
	shippingUseCase := usecase.NewShippingUseCase(shippingRepo)

	// Setup handlers
	shippingHandler := handler.NewShippingHandler(shippingUseCase)

	// Setup router
	r := gin.Default()

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app":        "ekspedisi-go",
			"version":    "1.0.0",
			"deployment": "development",
		})
	})

	// Health check endpoints
	r.GET("/live", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "alive"})
	})

	// Setup API group with auth middleware
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	// Shipping routes
	// Define level constants
	const (
		SUPER_ADMIN = 0
		MANAGER     = 100
		USER        = 1000
	)

	shipping := api.Group("/shipping")
	shipping.Use(middleware.RoleMiddleware(USER, MANAGER, SUPER_ADMIN))
	{
		shipping.GET("", shippingHandler.GetShipping)
		shipping.GET("/:id", shippingHandler.GetShippingByID)
	}

	// Start server
	serverAddr := ":" + cfg.ServerPort
	log.Printf("Starting server on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}