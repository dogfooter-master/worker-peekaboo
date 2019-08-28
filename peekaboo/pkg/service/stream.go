package service

import (
	"bytes"
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/pion/webrtc"
	"image/jpeg"
	"os"
	"time"
)
type StreamObject struct {
	Command string
	Quality int
	Fps int
	Handle w32.HWND
}
var StreamingRequest chan StreamObject

func init() {
	StreamingRequest = make(chan StreamObject, 1)
	go Streaming()
}

func Streaming() {
	var quality = 50
	var fps = 12
	var handle w32.HWND = 0
	for {
		select {
			default:
				if handle > 0 {
					fmt.Fprintf(os.Stderr, "handle: %v\n", handle)
					img := peekabooWindowInfo.GetWindowScreenShot(handle)
					if WebRTCMap != nil {
						//defer timeTrack(time.Now(), GetFunctionName() + "-22")
						for _, v := range WebRTCMap {
							//fmt.Fprintf(os.Stderr, "DEBUG: '%v'-'%v'\n", v.DataChannel.Label(), v.DataChannel.ReadyState())
							if v.DataChannel.ReadyState() == webrtc.DataChannelStateOpen {
								buf := new(bytes.Buffer)
								//_ = png.Encode(buf, img)
								_ = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
								s := buf.Bytes()
								//s = []byte{ 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10}

								//fmt.Fprintf(os.Stderr, "DEBUG: %v %v\n", len(s), math.MaxUint16)
								if err := SendImageBuffer(v.DataChannel, s); err != nil {
									fmt.Fprintf(os.Stderr, "%v\n", err)
								}
							}
						}
					}
				}
		case ss := <- StreamingRequest:
				switch ss.Command {
				case "start":
					handle = ss.Handle
				case "stop":
					handle = 0
				case "quality":
					quality = ss.Quality
					if quality > 100 {
						quality = 100
					} else if quality < 1 {
						quality = 1
					}
				case "fps":
					fps = ss.Fps
				}
		}
		if fps > 60 {
			fps = 60
		} else if fps < 5 {
			fps = 5
		}
		delay := time.Duration(1000 / fps)
		time.Sleep(delay * time.Millisecond)
	}
}