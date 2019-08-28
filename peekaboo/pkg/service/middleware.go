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
