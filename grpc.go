// +build grpc

package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"go.imgur.com/comments/handler"
)

const (
	// ServiceName is the name of this service.
	ServiceName = "comment"

	// ServerAddress is the address at which this service will run.
	ServerAddress = "127.0.0.1:8445"
)

func main() {
	// Create service
	srv := service.New(
		service.Name(ServiceName),
		service.Address(ServerAddress),
	)

	// Register handler
	srv.Handle(handler.NewComment())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
