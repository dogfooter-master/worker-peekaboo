package service

import (
	"context"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"

	log "github.com/go-kit/kit/log"
)

type Middleware func(PeekabooService) PeekabooService

type loggingMiddleware struct {
	logger log.Logger
	next   PeekabooService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PeekabooService) PeekabooService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Pikabu(ctx context.Context, req *pb.PikabuRequest) (res *pb.PikabuReply, err error) {
	defer func() {
		l.logger.Log("method", "Pikabu", "req", req, "res", res, "err", err)
	}()
	return l.next.Pikabu(ctx, req)
}

func (l loggingMiddleware) RefreshWindows(ctx context.Context, req *pb.RefreshWindowsRequest) (res *pb.RefreshWindowsReply, err error) {
	defer func() {
		l.logger.Log("method", "RefreshWindows", "req", req, "res", res, "err", err)
	}()
	return l.next.RefreshWindows(ctx, req)
}

func (l loggingMiddleware) StartStreaming(ctx context.Context, req *pb.StartStreamingRequest) (res *pb.StartStreamingReply, err error) {
	defer func() {
		l.logger.Log("method", "StartStreaming", "req", req, "res", res, "err", err)
	}()
	return l.next.StartStreaming(ctx, req)
}
func (l loggingMiddleware) EndStreaming(ctx context.Context, req *pb.EndStreamingRequest) (res *pb.EndStreamingReply, err error) {
	defer func() {
		l.logger.Log("method", "EndStreaming", "req", req, "res", res, "err", err)
	}()
	return l.next.EndStreaming(ctx, req)
}

func (l loggingMiddleware) ChangeQuality(ctx context.Context, req *pb.ChangeQualityRequest) (res *pb.ChangeQualityReply, err error) {
	defer func() {
		l.logger.Log("method", "ChangeQuality", "req", req, "res", res, "err", err)
	}()
	return l.next.ChangeQuality(ctx, req)
}
func (l loggingMiddleware) ChangeFps(ctx context.Context, req *pb.ChangeFpsRequest) (res *pb.ChangeFpsReply, err error) {
	defer func() {
		l.logger.Log("method", "ChangeFps", "req", req, "res", res, "err", err)
	}()
	return l.next.ChangeFps(ctx, req)
}

func (l loggingMiddleware) ChangeProperties(ctx context.Context, req *pb.ChangePropertiesRequest) (res *pb.ChangePropertiesReply, err error) {
	defer func() {
		l.logger.Log("method", "ChangeProperties", "req", req, "res", res, "err", err)
	}()
	return l.next.ChangeProperties(ctx, req)
}

func (l loggingMiddleware) MouseDown(ctx context.Context, req *pb.MouseDownRequest) (res *pb.MouseDownReply, err error) {
	defer func() {
		l.logger.Log("method", "MouseDown", "req", req, "res", res, "err", err)
	}()
	return l.next.MouseDown(ctx, req)
}

func (l loggingMiddleware) MouseDown2(ctx context.Context, req *pb.MouseDown2Request) (res *pb.MouseDown2Reply, err error) {
	defer func() {
		l.logger.Log("method", "MouseDown2", "req", req, "res", res, "err", err)
	}()
	return l.next.MouseDown2(ctx, req)
}

func (l loggingMiddleware) MouseUp(ctx context.Context, req *pb.MouseUpRequest) (res *pb.MouseUpReply, err error) {
	defer func() {
		l.logger.Log("method", "MouseUp", "req", req, "res", res, "err", err)
	}()
	return l.next.MouseUp(ctx, req)
}
func (l loggingMiddleware) MouseUp2(ctx context.Context, req *pb.MouseUp2Request) (res *pb.MouseUp2Reply, err error) {
	defer func() {
		l.logger.Log("method", "MouseUp2", "req", req, "res", res, "err", err)
	}()
	return l.next.MouseUp2(ctx, req)
}
func (l loggingMiddleware) MouseMove(ctx context.Context, req *pb.MouseMoveRequest) (res *pb.MouseMoveReply, err error) {
	defer func() {
		l.logger.Log("method", "MouseMove", "req", req, "res", res, "err", err)
	}()
	return l.next.MouseMove(ctx, req)
}
func (l loggingMiddleware) MouseMove2(ctx context.Context, req *pb.MouseMove2Request) (res *pb.MouseMove2Reply, err error) {
	defer func() {
		l.logger.Log("method", "MouseMove2", "req", req, "res", res, "err", err)
	}()
	return l.next.MouseMove2(ctx, req)
}
