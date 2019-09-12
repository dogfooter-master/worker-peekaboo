package service

import (
	"context"
	"fmt"
	"os"
	"time"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"

	"github.com/TheTitanrain/w32"
)

// PeekabooService describes the service.
type PeekabooService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Pikabu(ctx context.Context, req *pb.PikabuRequest) (res *pb.PikabuReply, err error)
	RefreshWindows(ctx context.Context, req *pb.RefreshWindowsRequest) (res *pb.RefreshWindowsReply, err error)
	StartStreaming(ctx context.Context, req *pb.StartStreamingRequest) (res *pb.StartStreamingReply, err error)
	EndStreaming(ctx context.Context, req *pb.EndStreamingRequest) (res *pb.EndStreamingReply, err error)
	ChangeQuality(ctx context.Context, req *pb.ChangeQualityRequest) (res *pb.ChangeQualityReply, err error)
	ChangeFps(ctx context.Context, req *pb.ChangeFpsRequest) (res *pb.ChangeFpsReply, err error)
	ChangeProperties(ctx context.Context, req *pb.ChangePropertiesRequest) (res *pb.ChangePropertiesReply, err error)
	MouseDown(ctx context.Context, req *pb.MouseDownRequest) (res *pb.MouseDownReply, err error)
	MouseDown2(ctx context.Context, req *pb.MouseDown2Request) (res *pb.MouseDown2Reply, err error)
	MouseUp(ctx context.Context, req *pb.MouseUpRequest) (res *pb.MouseUpReply, err error)
	MouseUp2(ctx context.Context, req *pb.MouseUp2Request) (res *pb.MouseUp2Reply, err error)
	MouseMove(ctx context.Context, req *pb.MouseMoveRequest) (res *pb.MouseMoveReply, err error)
	MouseMove2(ctx context.Context, req *pb.MouseMove2Request) (res *pb.MouseMove2Reply, err error)
}

type basicPeekabooService struct{}

func (b *basicPeekabooService) Pikabu(ctx context.Context, req *pb.PikabuRequest) (res *pb.PikabuReply, err error) {

	//defer timeTrack(time.Now(), GetFunctionName())
	//fmt.Fprintf(os.Stderr, "DEBUG: Category: %v\n", req.Category)

	res = &pb.PikabuReply{
		Category: "response_" + req.Category,
	}
	return res, err
}

func (b *basicPeekabooService) RefreshWindows(ctx context.Context, req *pb.RefreshWindowsRequest) (res *pb.RefreshWindowsReply, err error) {
	defer timeTrack(time.Now(), GetFunctionName())
	peekabooWindowInfo = PeekabooWin{
		Wildcard: req.Keyword,
	}

	peekabooWindowInfo.FindWindowWildcard2()

	var windowList []*pb.RefreshWindowsReply_Window
	for k, v := range peekabooWindowInfo.HandleMap {
		windowList = append(windowList,
			&pb.RefreshWindowsReply_Window{
				Title:  v.Title,
				Handle: int32(k),
			})
	}

	res = &pb.RefreshWindowsReply{
		WindowList: windowList,
	}
	return res, err
}

func (b *basicPeekabooService) StartStreaming(ctx context.Context, req *pb.StartStreamingRequest) (res *pb.StartStreamingReply, err error) {
	defer timeTrack(time.Now(), GetFunctionName())
	fmt.Fprintf(os.Stderr, "DEBUG: %#v\n", req)
	ss := StreamObject{
		Command:      "start",
		Handle:       w32.HWND(req.Handle),
		ChannelLabel: req.Label,
	}
	if req.Fps > 0 {
		ss.Fps = int(req.Fps)
	} else {
		ss.Fps = 24
	}
	if req.Quality > 0 {
		ss.Quality = int(req.Quality)
	} else {
		ss.Quality = 50
	}
	StreamingRequest <- ss
	res = &pb.StartStreamingReply{
		Label: req.Label,
	}
	return res, err
}
func (b *basicPeekabooService) EndStreaming(ctx context.Context, req *pb.EndStreamingRequest) (res *pb.EndStreamingReply, err error) {
	StreamingRequest <- StreamObject{
		Command:      "stop",
		ChannelLabel: req.Label,
	}

	res = &pb.EndStreamingReply{
		Handle: req.Handle,
	}
	return res, err
}
func (b *basicPeekabooService) ChangeQuality(ctx context.Context, req *pb.ChangeQualityRequest) (res *pb.ChangeQualityReply, err error) {
	StreamingRequest <- StreamObject{
		Command:      "quality",
		ChannelLabel: req.Label,
		Quality:      int(req.Quality),
	}

	res = &pb.ChangeQualityReply{
		Quality: req.Quality,
	}
	return res, err
}
func (b *basicPeekabooService) ChangeFps(ctx context.Context, req *pb.ChangeFpsRequest) (res *pb.ChangeFpsReply, err error) {
	StreamingRequest <- StreamObject{
		Command:      "fps",
		ChannelLabel: req.Label,
		Fps:          int(req.Fps),
	}

	res = &pb.ChangeFpsReply{
		Fps: req.Fps,
	}
	return res, err
}

func (b *basicPeekabooService) ChangeProperties(ctx context.Context, req *pb.ChangePropertiesRequest) (res *pb.ChangePropertiesReply, err error) {
	StreamingRequest <- StreamObject{
		Command:      "change",
		Handle:       w32.HWND(req.Handle),
		ChannelLabel: req.Label,
		Fps:          int(req.Fps),
		Quality:      int(req.Quality),
	}

	res = &pb.ChangePropertiesReply{
		Label: req.Label,
	}
	return res, err
}

func (b *basicPeekabooService) MouseDown(ctx context.Context, req *pb.MouseDownRequest) (res *pb.MouseDownReply, err error) {
	peekabooWindowInfo.MouseDown(w32.HWND(req.Handle), req.X, req.Y)

	res = &pb.MouseDownReply{
		Handle: req.Handle,
	}
	return res, err
}
func (b *basicPeekabooService) MouseDown2(ctx context.Context, req *pb.MouseDown2Request) (res *pb.MouseDown2Reply, err error) {
	peekabooWindowInfo.MouseDown2(w32.HWND(req.Handle), int(req.X), int(req.Y))

	res = &pb.MouseDown2Reply{
		Handle: req.Handle,
	}
	return res, err
}

func (b *basicPeekabooService) MouseUp(ctx context.Context, req *pb.MouseUpRequest) (res *pb.MouseUpReply, err error) {
	peekabooWindowInfo.MouseUp(w32.HWND(req.Handle), req.X, req.Y)

	res = &pb.MouseUpReply{
		Handle: req.Handle,
	}
	return res, err
}
func (b *basicPeekabooService) MouseUp2(ctx context.Context, req *pb.MouseUp2Request) (res *pb.MouseUp2Reply, err error) {
	peekabooWindowInfo.MouseUp2(w32.HWND(req.Handle), int(req.X), int(req.Y))

	res = &pb.MouseUp2Reply{
		Handle: req.Handle,
	}
	return res, err
}
func (b *basicPeekabooService) MouseMove(ctx context.Context, req *pb.MouseMoveRequest) (res *pb.MouseMoveReply, err error) {

	peekabooWindowInfo.MouseMove(w32.HWND(req.Handle), req.X, req.Y)

	res = &pb.MouseMoveReply{
		Handle: req.Handle,
	}
	return res, err
}
func (b *basicPeekabooService) MouseMove2(ctx context.Context, req *pb.MouseMove2Request) (res *pb.MouseMove2Reply, err error) {

	peekabooWindowInfo.MouseMove2(w32.HWND(req.Handle), int(req.X), int(req.Y))

	res = &pb.MouseMove2Reply{
		Handle: req.Handle,
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
