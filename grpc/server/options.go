package grpcserver

import "github.com/gaiaz-iusipov/go-app/grpc"

type Option func(s *Server)

func WithService(service grpc.Service) Option {
	return func(s *Server) {
		s.RegisterService(service.Desc(), service.Impl())
	}
}
