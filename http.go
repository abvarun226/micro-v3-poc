// +build http

package main

import (
	"log"
	"time"

	"github.com/go-chi/chi"

	"go.imgur.com/comments/handler"
	"go.imgur.com/comments/register"
	"go.imgur.com/comments/server"
)

const (
	// ServiceName is the name of this service.
	ServiceName = "comment"

	// ServerAddress is the address at which this service will run.
	ServerAddress = "127.0.0.1:8445"
)

func main() {
	reg := register.New(
		register.WithName(ServiceName),
		register.WithTTL(10*time.Second),
		register.WithInterval(10*time.Second),
		register.WithAddr(ServerAddress),
	)

	// Set up our HTTP routes.
	r := chi.NewRouter()
	r.Route("/comment/v1", func(r chi.Router) {
		r.Get("/hello", handler.HelloWorld)
	})

	reg.Register()

	srv := server.New(
		server.WithHandler(r),
		server.WithAddr(ServerAddress),
		server.WithShutdownGracePeriod(10*time.Second),
	)
	srv.Start()

	log.Printf("server successfully shutdown")
}
