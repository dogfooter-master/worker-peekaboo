package grpc

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	"worker-peekaboo/peekaboo/pkg/endpoint"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"
)

// makePikabuHandler creates the handler logic
func makePikabuHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.PikabuEndpoint, decodePikabuRequest, encodePikabuResponse, options...)
}

// decodePikabuResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodePikabuRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PikabuRequest)
	return endpoint.PikabuRequest{
		Req: req,
	}, nil
}

// encodePikabuResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodePikabuResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.PikabuResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) Pikabu(ctx context1.Context, req *pb.PikabuRequest) (*pb.PikabuReply, error) {
	_, rep, err := g.pikabu.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PikabuReply), nil
}
