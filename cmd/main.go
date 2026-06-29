package main

import (
	"log"
	"time"

	"vibrox-core/internal/config"
	"vibrox-core/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	router := gin.New()

	time.Sleep(5 * time.Second)

	config.Connect()

	logClientConn, err := config.InitLoggerClient()
	if err != nil {
		log.Fatal("Failed to initialize logger client:", err)
	}

	authClientConn, err := config.InitAuthClient()
	if err != nil {
		log.Fatal("Failed to initialize auth client:", err)
	}

	defer func() {
		if err := logClientConn.Close(); err != nil {
			log.Fatal("Failed to close log client:", err)
		}
	}()
	defer func() {
		if err = authClientConn.Close(); err != nil {
			log.Fatal("Failed to close auth client:", err)
		}
	}()

	routes.UserRoute(router)
	if err := router.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
