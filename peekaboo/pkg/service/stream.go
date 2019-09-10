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
	Command      string
	Quality      int
	Fps          int
	Handle       w32.HWND
	ChannelLabel string
	Err          string
}

var StreamingRequest chan StreamObject
var StreamingResponse chan StreamObject
var StreamingKeepGoing chan StreamObject
var CurrentStream StreamObject

func init() {
	StreamingRequest = make(chan StreamObject, 1)
	StreamingResponse = make(chan StreamObject, 1)
	StreamingKeepGoing = make(chan StreamObject, 1)
	for i := 0; i < 1; i++ {
		go Streaming(i)
	}
}

func Streaming(worker int) {
	var ss StreamObject
	for {
		select {
		case ss = <-StreamingRequest:
			s := ss
			for len(StreamingKeepGoing) > 0 {
				s = <-StreamingKeepGoing
			}
			s.Command = ss.Command
			switch ss.Command {
			case "start":
				StreamingKeepGoing <- ss
			case "stop":
				StreamingKeepGoing <- ss
			case "quality":
				StreamingKeepGoing <- ss
			case "fps":
				StreamingKeepGoing <- ss
			case "change":
				StreamingKeepGoing <- ss
			}
			fmt.Fprintf(os.Stderr, "DEBUG: '%v'-'%#v'\n", worker, ss)
		case ss = <- StreamingKeepGoing:
			if ss.Fps > 60 {
				ss.Fps = 60
			} else if ss.Fps < 5 {
				ss.Fps = 5
			}
			delay := time.Duration(1000 / ss.Fps)
			time.Sleep(delay * time.Millisecond)

			if ss.Handle > 0 {
				img := peekabooWindowInfo.GetWindowScreenShot(ss.Handle)
				isFound := false
				if img != nil {
					if WebRTCMap != nil {
						//defer timeTrack(time.Now(), GetFunctionName() + "-22")
						if v, ok := WebRTCMap[ss.ChannelLabel]; ok {
							//fmt.Fprintf(os.Stderr, "DEBUG: '%v'-'%v'-'%v'-'%v'-'%v'-'%v'\n", worker, v.DataChannel.Label(), ss.Handle, ss.Fps, ss.Quality, img.Bounds())
							if v.DataChannel.ReadyState() == webrtc.DataChannelStateOpen && v.DataChannel.Label() == ss.ChannelLabel {
								isFound = true
								buf := new(bytes.Buffer)
								//_ = png.Encode(buf, img)
								if ss.Quality < 1 {
									ss.Quality = 1
								} else if ss.Quality > 100 {
									ss.Quality = 100
								}
								_ = jpeg.Encode(buf, img, &jpeg.Options{Quality: ss.Quality})
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
				if isFound == false || ss.Command == "stop" {
					fmt.Println("not found #", ss.ChannelLabel)
					CurrentStream = StreamObject{}
				} else {
					CurrentStream = ss
					StreamingKeepGoing <- ss
				}
			}
			fmt.Fprintf(os.Stderr, "StreamingKeepGoing\n")
		default:
			fmt.Fprintf(os.Stderr, "Default\n")
			time.Sleep(1 * time.Second)
		}
	}
}
