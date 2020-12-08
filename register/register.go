package register

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/client"
	grpcCli "github.com/micro/micro/v3/service/client/grpc"
	"github.com/micro/micro/v3/service/registry"
	registrySrv "github.com/micro/micro/v3/service/registry/client"
)

type service struct {
	Registry registry.Registry
	opts     Options
}

// Service interface.
type Service interface {
	Register()
}

// New service registry.
func New(opt ...Option) Service {
	client.DefaultClient = grpcCli.NewClient()

	opts := NewOptions(opt...)

	return &service{
		Registry: registrySrv.NewRegistry(),
		opts:     opts,
	}
}

// Register will register a given service with micro.
func (s *service) Register() {
	id, _ := uuid.NewUUID()
	nodeID := s.opts.Name + "-" + id.String()

	svc := &registry.Service{
		Name: s.opts.Name,
		Nodes: []*registry.Node{
			{
				Id:      nodeID,
				Address: s.opts.Addr,
			},
		},
	}

	log.Printf("[INFO] registering service: %s", nodeID)
	s.Registry.Register(svc, registry.RegisterTTL(s.opts.TTL))

	go func() {
		for range time.Tick(s.opts.Interval) {
			s.Registry.Register(svc, registry.RegisterTTL(s.opts.TTL))
		}
	}()
}
