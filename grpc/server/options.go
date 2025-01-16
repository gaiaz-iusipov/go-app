package grpcserver

import (
	"github.com/gaiaz-iusipov/go-app/grpc"
	"google.golang.org/grpc/reflection"
)

type Option func(s *Server)

func WithService(service grpc.Service) Option {
	return func(s *Server) {
		s.RegisterService(service.Desc(), service.Impl())
	}
}

func EnableReflection() Option {
	return func(s *Server) {
		reflection.Register(s.grpcServer)
	}
}
