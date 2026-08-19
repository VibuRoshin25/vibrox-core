package main

import (
	"log"
	"os"

	"vibrox-core/internal/config"
	"vibrox-core/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using environment variables")
	}

	loggerConnection, err := config.InitLoggerClient()
	if err != nil {
		log.Fatal("initialize logger client: ", err)
	}
	defer func() {
		if err := loggerConnection.Close(); err != nil {
			log.Printf("close logger client: %v", err)
		}
	}()
	arenaConnection, err := config.InitArenaClient()
	if err != nil {
		log.Fatal("initialize arena client: ", err)
	}
	defer func() {
		if err := arenaConnection.Close(); err != nil {
			log.Printf("close arena client: %v", err)
		}
	}()

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	routes.Register(router)

	address := ":" + envOrDefault("PORT", "8080")
	if err := router.Run(address); err != nil {
		log.Fatal("start server: ", err)
	}
}

func envOrDefault(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
