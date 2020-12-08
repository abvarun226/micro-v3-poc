package register

import (
	"time"
)

const (
	// DefaultName is the default service name.
	DefaultName = "imgur"

	// DefaultAddr is the default service address.
	DefaultAddr = "127.0.0.1:9090"

	// DefaultTTL is the default service registration TTL.
	DefaultTTL = 10 * time.Second

	// DefaultInterval is the default service registration interval.
	DefaultInterval = 10 * time.Second
)

// Options ...
type Options struct {
	Name     string
	Addr     string
	TTL      time.Duration
	Interval time.Duration
}

// Option ...
type Option func(*Options)

// NewOptions returns options for register struct.
func NewOptions(options ...Option) Options {
	opts := Options{
		Name:     DefaultName,
		Addr:     DefaultAddr,
		TTL:      DefaultTTL,
		Interval: DefaultInterval,
	}

	for _, o := range options {
		o(&opts)
	}

	return opts
}

// WithAddr sets the service address.
func WithAddr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

// WithTTL sets the service registration TTL.
func WithTTL(ttl time.Duration) Option {
	return func(o *Options) {
		o.TTL = ttl
	}
}

// WithInterval sets the service registration interval.
func WithInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.Interval = interval
	}
}

// WithName sets the service name.
func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}
