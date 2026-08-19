package config

import (
	"os"

	arenapb "vibrox-core/internal/proto/arena"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ArenaClient is the gRPC client for the Arena service.
var ArenaClient arenapb.ArenaClient

// InitArenaClient initializes the Arena client.
func InitArenaClient() (*grpc.ClientConn, error) {
	host := os.Getenv("ARENA_HOST")
	if host == "" {
		host = "arena:8100"
	}
	connection, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	ArenaClient = arenapb.NewArenaClient(connection)
	return connection, nil
}
