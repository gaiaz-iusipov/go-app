package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	grpcpkg "github.com/gaiaz-iusipov/go-app/grpc"
)

var _ grpcpkg.Server = (*Server)(nil)

func New(addr string, opts ...Option) Server {
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	server := Server{
		addr:       addr,
		grpcServer: grpcServer,
	}
	for _, opt := range opts {
		opt(&server)
	}
	return server
}

type Server struct {
	addr       string
	grpcServer *grpc.Server
}

func (s Server) RegisterService(desc *grpc.ServiceDesc, impl any) {
	s.grpcServer.RegisterService(desc, impl)
}

func (s Server) Run(ctx context.Context) error {
	listener, err := new(net.ListenConfig).Listen(ctx, "tcp", s.addr)
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}

	go func() {
		if serveErr := s.grpcServer.Serve(listener); !errors.Is(serveErr, grpc.ErrServerStopped) {
			slog.ErrorContext(ctx, "failed to serve grpc server", "error", serveErr)
		}
	}()

	return nil
}

func (s Server) GracefulStop() {
	s.grpcServer.GracefulStop()
}
