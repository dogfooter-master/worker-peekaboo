package service

import (
	"fmt"
	"github.com/pion/webrtc"
	"math"
	"os"
)

func SendImageBuffer(dc *webrtc.DataChannel, buffer []byte) (err error) {
	//defer timeTrack(time.Now(), GetFunctionName())
	if err = dc.SendText("i-s"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	if err = SendImageBufferRecursive(dc, buffer); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	if err = dc.SendText("i-e"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	return
}

func SendImageBufferRecursive(dc *webrtc.DataChannel, buffer []byte) (err error) {

	var s []byte
	var remain []byte
	if len(buffer) > math.MaxUint16 {
		s = buffer[:math.MaxUint16 - 1]
		remain = buffer[math.MaxUint16 - 1:]
	} else {
		s = buffer
	}
	sendErr := dc.Send(s)
	if sendErr != nil {
		fmt.Fprintf(os.Stderr, "DEBUG: %v\n", sendErr)
	}

	if len(remain) > 0 {
		if err = SendImageBufferRecursive(dc, remain); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
	}

	return
}
