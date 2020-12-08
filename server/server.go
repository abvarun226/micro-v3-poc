package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	opts Options
}

// Server interface.
type Server interface {
	Start()
}

// New http server.
func New(opt ...Option) Server {
	opts := NewOptions(opt...)
	return &server{
		opts: opts,
	}
}

// Start the server.
func (s *server) Start() {
	srv := &http.Server{
		Addr:    s.opts.Addr,
		Handler: s.opts.Handler,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] listen: %v", err)
		}
	}()

	log.Printf("[INFO] server started on address: %s", srv.Addr)

	<-done

	log.Printf("[INFO] server shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), s.opts.ShutdownGracePeriod)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] server failed to shutdown: %v", err)
	}
}
