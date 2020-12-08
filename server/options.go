package server

import (
	"net/http"
	"time"
)

const (
	// DefaultAddr is the default server address.
	DefaultAddr = "127.0.0.1:9090"

	// DefaultShutdownGracePeriod is the default server shutdown grace period.
	DefaultShutdownGracePeriod = 10 * time.Second
)

// Options ...
type Options struct {
	Handler             http.Handler
	Addr                string
	ShutdownGracePeriod time.Duration
}

// Option ...
type Option func(*Options)

// NewOptions returns options for register struct.
func NewOptions(options ...Option) Options {
	opts := Options{
		Addr:                DefaultAddr,
		ShutdownGracePeriod: DefaultShutdownGracePeriod,
	}

	for _, o := range options {
		o(&opts)
	}

	return opts
}

// WithHandler sets the server handler.
func WithHandler(h http.Handler) Option {
	return func(o *Options) {
		o.Handler = h
	}
}

// WithAddr sets the server address.
func WithAddr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

// WithShutdownGracePeriod sets the server shutdown grace period.
func WithShutdownGracePeriod(period time.Duration) Option {
	return func(o *Options) {
		o.ShutdownGracePeriod = period
	}
}
