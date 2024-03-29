// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "worker-peekaboo/peekaboo/pkg/endpoint"
	pb "worker-peekaboo/peekaboo/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	pikabu           grpc.Handler
	refreshWindows   grpc.Handler
	startStreaming   grpc.Handler
	endStreaming     grpc.Handler
	changeQuality    grpc.Handler
	changeFps        grpc.Handler
	changeProperties grpc.Handler
	mouseDown        grpc.Handler
	mouseDown2       grpc.Handler
	mouseUp          grpc.Handler
	mouseUp2         grpc.Handler
	mouseMove        grpc.Handler
	mouseMove2       grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.PeekabooServer {
	return &grpcServer{
		changeFps:        makeChangeFpsHandler(endpoints, options["ChangeFps"]),
		changeProperties: makeChangePropertiesHandler(endpoints, options["ChangeProperties"]),
		changeQuality:    makeChangeQualityHandler(endpoints, options["ChangeQuality"]),
		endStreaming:     makeEndStreamingHandler(endpoints, options["EndStreaming"]),
		mouseDown:        makeMouseDownHandler(endpoints, options["MouseDown"]),
		mouseDown2:       makeMouseDown2Handler(endpoints, options["MouseDown2"]),
		mouseMove:        makeMouseMoveHandler(endpoints, options["MouseMove"]),
		mouseMove2:       makeMouseMove2Handler(endpoints, options["MouseMove2"]),
		mouseUp:          makeMouseUpHandler(endpoints, options["MouseUp"]),
		mouseUp2:         makeMouseUp2Handler(endpoints, options["MouseUp2"]),
		pikabu:           makePikabuHandler(endpoints, options["Pikabu"]),
		refreshWindows:   makeRefreshWindowsHandler(endpoints, options["RefreshWindows"]),
		startStreaming:   makeStartStreamingHandler(endpoints, options["StartStreaming"]),
	}
}
