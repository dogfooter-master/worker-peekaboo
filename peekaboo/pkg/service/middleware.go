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
