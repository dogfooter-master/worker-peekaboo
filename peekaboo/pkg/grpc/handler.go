package grpc

import (
	"context"
	"worker-peekaboo/peekaboo/pkg/endpoint"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"

	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

func makePikabuHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.PikabuEndpoint, decodePikabuRequest, encodePikabuResponse, options...)
}

func decodePikabuRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PikabuRequest)
	return endpoint.PikabuRequest{
		Req: req,
	}, nil
}

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

func makeRefreshWindowsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RefreshWindowsEndpoint, decodeRefreshWindowsRequest, encodeRefreshWindowsResponse, options...)
}

func decodeRefreshWindowsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.RefreshWindowsRequest)
	return endpoint.RefreshWindowsRequest{
		Req: req,
	}, nil
}

func encodeRefreshWindowsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.RefreshWindowsResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) RefreshWindows(ctx context1.Context, req *pb.RefreshWindowsRequest) (*pb.RefreshWindowsReply, error) {
	_, rep, err := g.refreshWindows.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RefreshWindowsReply), nil
}

func makeStartStreamingHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.StartStreamingEndpoint, decodeStartStreamingRequest, encodeStartStreamingResponse, options...)
}

func decodeStartStreamingRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.StartStreamingRequest)
	return endpoint.StartStreamingRequest{
		Req: req,
	}, nil
}

func encodeStartStreamingResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.StartStreamingResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) StartStreaming(ctx context1.Context, req *pb.StartStreamingRequest) (*pb.StartStreamingReply, error) {
	_, rep, err := g.startStreaming.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.StartStreamingReply), nil
}

func makeEndStreamingHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.EndStreamingEndpoint, decodeEndStreamingRequest, encodeEndStreamingResponse, options...)
}

func decodeEndStreamingRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.EndStreamingRequest)
	return endpoint.EndStreamingRequest{
		Req: req,
	}, nil
}

func encodeEndStreamingResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.EndStreamingResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) EndStreaming(ctx context1.Context, req *pb.EndStreamingRequest) (*pb.EndStreamingReply, error) {
	_, rep, err := g.endStreaming.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.EndStreamingReply), nil
}

func makeChangeQualityHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ChangeQualityEndpoint, decodeChangeQualityRequest, encodeChangeQualityResponse, options...)
}

func decodeChangeQualityRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ChangeQualityRequest)
	return endpoint.ChangeQualityRequest{
		Req: req,
	}, nil
}

func encodeChangeQualityResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ChangeQualityResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) ChangeQuality(ctx context1.Context, req *pb.ChangeQualityRequest) (*pb.ChangeQualityReply, error) {
	_, rep, err := g.changeQuality.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ChangeQualityReply), nil
}

func makeChangeFpsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ChangeFpsEndpoint, decodeChangeFpsRequest, encodeChangeFpsResponse, options...)
}

func decodeChangeFpsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ChangeFpsRequest)
	return endpoint.ChangeFpsRequest{
		Req: req,
	}, nil
}

func encodeChangeFpsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ChangeFpsResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) ChangeFps(ctx context1.Context, req *pb.ChangeFpsRequest) (*pb.ChangeFpsReply, error) {
	_, rep, err := g.changeFps.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ChangeFpsReply), nil
}
