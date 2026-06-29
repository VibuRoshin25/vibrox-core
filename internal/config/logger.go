package config

import (
	"os"

	"vibrox-core/internal/proto/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var loggerHost = os.Getenv("LOGGER_HOST")

// LogClient is the logging interface
var LogClient logger.LoggerClient

// InitLoggerClient initialize logger client
func InitLoggerClient() (*grpc.ClientConn, error) {

	conn, err := grpc.NewClient(loggerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	LogClient = logger.NewLoggerClient(conn)

	return conn, nil
}
