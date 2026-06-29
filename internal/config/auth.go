package config

import (
	"os"

	"vibrox-core/internal/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var authHost = os.Getenv("AUTH_HOST")

// AuthClient is the authentication interface
var AuthClient auth.TokenClient

// InitAuthClient initializes authentication client
func InitAuthClient() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(authHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	AuthClient = auth.NewTokenClient(conn)

	return conn, nil
}
