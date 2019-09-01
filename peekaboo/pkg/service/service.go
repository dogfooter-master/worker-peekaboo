package service

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"os"
	"time"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"

	"github.com/TheTitanrain/w32"

	"github.com/pion/webrtc"
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
}

type basicPeekabooService struct{}

func (b *basicPeekabooService) Pikabu(ctx context.Context, req *pb.PikabuRequest) (res *pb.PikabuReply, err error) {

	//defer timeTrack(time.Now(), GetFunctionName())
	//fmt.Fprintf(os.Stderr, "DEBUG: Category: %v\n", req.Category)

	pWin := PeekabooWin{
		Wildcard: req.Keyword,
	}
	pWin.FindWindowWildcard()

	//fmt.Println("Result:", pWin)

	if len(pWin.HWNDList) > 0 {
		img := pWin.GetWindowScreenShot(pWin.HWNDList[0])
		if WebRTCMap != nil {
			//defer timeTrack(time.Now(), GetFunctionName() + "-22")
			for _, v := range WebRTCMap {
				//fmt.Fprintf(os.Stderr, "DEBUG: '%v'-'%v'\n", v.DataChannel.Label(), v.DataChannel.ReadyState())
				if v.DataChannel.ReadyState() == webrtc.DataChannelStateOpen {
					buf := new(bytes.Buffer)
					//_ = png.Encode(buf, img)
					_ = jpeg.Encode(buf, img, &jpeg.Options{Quality: 75})
					s := buf.Bytes()
					//s = []byte{ 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10}

					//fmt.Fprintf(os.Stderr, "DEBUG: %v %v\n", len(s), math.MaxUint16)
					if err = SendImageBuffer(v.DataChannel, s); err != nil {
						return
					}
				}
			}
		}
	}

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

	peekabooWindowInfo.FindWindowWildcard()

	var windowList []*pb.RefreshWindowsReply_Window
	for _, e := range peekabooWindowInfo.HWNDList {
		windowList = append(windowList,
			&pb.RefreshWindowsReply_Window{
				Title:  peekabooWindowInfo.TitleMap[e],
				Handle: int32(e),
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
