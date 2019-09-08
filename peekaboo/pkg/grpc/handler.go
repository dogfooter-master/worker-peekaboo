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

func makeChangePropertiesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ChangePropertiesEndpoint, decodeChangePropertiesRequest, encodeChangePropertiesResponse, options...)
}

func decodeChangePropertiesRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ChangePropertiesRequest)
	return endpoint.ChangePropertiesRequest{
		Req: req,
	}, nil
}

func encodeChangePropertiesResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.ChangePropertiesResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) ChangeProperties(ctx context1.Context, req *pb.ChangePropertiesRequest) (*pb.ChangePropertiesReply, error) {
	_, rep, err := g.changeProperties.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ChangePropertiesReply), nil
}

func makeMouseDownHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.MouseDownEndpoint, decodeMouseDownRequest, encodeMouseDownResponse, options...)
}

func decodeMouseDownRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.MouseDownRequest)
	return endpoint.MouseDownRequest{
		Req: req,
	}, nil
}

func encodeMouseDownResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.MouseDownResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) MouseDown(ctx context1.Context, req *pb.MouseDownRequest) (*pb.MouseDownReply, error) {
	_, rep, err := g.mouseDown.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.MouseDownReply), nil
}

func makeMouseDown2Handler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.MouseDown2Endpoint, decodeMouseDown2Request, encodeMouseDown2Response, options...)
}

func decodeMouseDown2Request(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.MouseDown2Request)
	return endpoint.MouseDown2Request{
		Req: req,
	}, nil
}

func encodeMouseDown2Response(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.MouseDown2Response)
	return res.Res, res.Failed()
}
func (g *grpcServer) MouseDown2(ctx context1.Context, req *pb.MouseDown2Request) (*pb.MouseDown2Reply, error) {
	_, rep, err := g.mouseDown2.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.MouseDown2Reply), nil
}

func makeMouseUpHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.MouseUpEndpoint, decodeMouseUpRequest, encodeMouseUpResponse, options...)
}

func decodeMouseUpRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.MouseUpRequest)
	return endpoint.MouseUpRequest{
		Req: req,
	}, nil
}

func encodeMouseUpResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.MouseUpResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) MouseUp(ctx context1.Context, req *pb.MouseUpRequest) (*pb.MouseUpReply, error) {
	_, rep, err := g.mouseUp.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.MouseUpReply), nil
}

func makeMouseUp2Handler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.MouseUp2Endpoint, decodeMouseUp2Request, encodeMouseUp2Response, options...)
}

func decodeMouseUp2Request(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.MouseUp2Request)
	return endpoint.MouseUp2Request{
		Req: req,
	}, nil
}

func encodeMouseUp2Response(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.MouseUp2Response)
	return res.Res, res.Failed()
}
func (g *grpcServer) MouseUp2(ctx context1.Context, req *pb.MouseUp2Request) (*pb.MouseUp2Reply, error) {
	_, rep, err := g.mouseUp2.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.MouseUp2Reply), nil
}

func makeMouseMoveHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.MouseMoveEndpoint, decodeMouseMoveRequest, encodeMouseMoveResponse, options...)
}

func decodeMouseMoveRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.MouseMoveRequest)
	return endpoint.MouseMoveRequest{
		Req: req,
	}, nil
}

func encodeMouseMoveResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.MouseMoveResponse)
	return res.Res, res.Failed()
}
func (g *grpcServer) MouseMove(ctx context1.Context, req *pb.MouseMoveRequest) (*pb.MouseMoveReply, error) {
	_, rep, err := g.mouseMove.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.MouseMoveReply), nil
}

func makeMouseMove2Handler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.MouseMove2Endpoint, decodeMouseMove2Request, encodeMouseMove2Response, options...)
}

func decodeMouseMove2Request(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.MouseMove2Request)
	return endpoint.MouseMove2Request{
		Req: req,
	}, nil
}

func encodeMouseMove2Response(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.MouseMove2Response)
	return res.Res, res.Failed()
}
func (g *grpcServer) MouseMove2(ctx context1.Context, req *pb.MouseMove2Request) (*pb.MouseMove2Reply, error) {
	_, rep, err := g.mouseMove2.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.MouseMove2Reply), nil
}
