package service

import (
	"context"
	"fmt"
	"os"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"
)

// PeekabooService describes the service.
type PeekabooService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Pikabu(ctx context.Context, req *pb.PikabuRequest) (res *pb.PikabuReply, err error)
}

type basicPeekabooService struct{}

func (b *basicPeekabooService) Pikabu(ctx context.Context, req *pb.PikabuRequest) (res *pb.PikabuReply, err error) {

	fmt.Fprintf(os.Stderr, "DEBUG: %v\n", req.Category)

	res = &pb.PikabuReply{
		Category: "response_" + req.Category,
	}
	return res, err
}

// NewBasicPeekabooService returns a naive, stateless implementation of PeekabooService.
func NewBasicPeekabooService() PeekabooService {
	return &basicPeekabooService{}
}

// New returns a PeekabooService with all of the expected middleware wired in.
func New(middleware []Middleware) PeekabooService {
	var svc PeekabooService = NewBasicPeekabooService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
