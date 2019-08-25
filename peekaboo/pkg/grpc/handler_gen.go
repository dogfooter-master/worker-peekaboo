// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "worker-peekaboo/peekaboo/pkg/endpoint"
	pb "worker-peekaboo/peekaboo/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	pikabu grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.PeekabooServer {
	return &grpcServer{pikabu: makePikabuHandler(endpoints, options["Pikabu"])}
}