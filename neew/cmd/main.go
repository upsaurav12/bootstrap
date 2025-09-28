package main

import (
	"log"
	"neew/internal/routers"
	"neew/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.New()

	// Setup Gin
	r := gin.Default()

	// Register routes
	api.RegisterRoutes(r)

	// Start server
	log.Printf("Server starting on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
