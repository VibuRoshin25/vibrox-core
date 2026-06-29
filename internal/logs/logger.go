package logs

import (
	"context"
	"log"

	"vibrox-core/internal/config"
	"vibrox-core/internal/proto/logger"
)

// LogError is used to log error messages
func LogError(ctx context.Context, message string) {

	resp, err := config.LogClient.Log(ctx, &logger.LogRequest{
		Message: message,
		Service: "core",
		Level:   "ERROR",
	})
	if err != nil {
		log.Println("Failed to log error: ", err)
	} else if !resp.Success {
		log.Println("Logger error: ", resp.Err)
	}
}

// LogInfo is used to log success messages
func LogInfo(ctx context.Context, message string) {
	resp, err := config.LogClient.Log(ctx, &logger.LogRequest{
		Message: message,
		Service: "core",
		Level:   "INFO",
	})
	if err != nil {
		log.Println("Failed to log error: ", err)
	} else if !resp.Success {
		log.Println("Logger error: ", resp.Err)
	}
}
