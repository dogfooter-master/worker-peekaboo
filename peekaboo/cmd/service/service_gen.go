// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "worker-peekaboo/peekaboo/pkg/endpoint"
	http1 "worker-peekaboo/peekaboo/pkg/http"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	initGRPCHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"ChangeFps":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ChangeFps", logger))},
		"ChangeProperties": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ChangeProperties", logger))},
		"ChangeQuality":    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ChangeQuality", logger))},
		"EndStreaming":     {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "EndStreaming", logger))},
		"MouseDown":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "MouseDown", logger))},
		"MouseDown2":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "MouseDown2", logger))},
		"MouseMove":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "MouseMove", logger))},
		"MouseMove2":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "MouseMove2", logger))},
		"MouseUp":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "MouseUp", logger))},
		"MouseUp2":         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "MouseUp2", logger))},
		"Pikabu":           {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Pikabu", logger))},
		"RefreshWindows":   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "RefreshWindows", logger))},
		"StartStreaming":   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "StartStreaming", logger))},
	}
	return options
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"ChangeFps":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ChangeFps", logger))},
		"ChangeProperties": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ChangeProperties", logger))},
		"ChangeQuality":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ChangeQuality", logger))},
		"EndStreaming":     {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "EndStreaming", logger))},
		"MouseDown":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "MouseDown", logger))},
		"MouseDown2":       {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "MouseDown2", logger))},
		"MouseMove":        {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "MouseMove", logger))},
		"MouseMove2":       {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "MouseMove2", logger))},
		"MouseUp":          {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "MouseUp", logger))},
		"MouseUp2":         {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "MouseUp2", logger))},
		"Pikabu":           {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Pikabu", logger))},
		"RefreshWindows":   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "RefreshWindows", logger))},
		"StartStreaming":   {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "StartStreaming", logger))},
	}
	return options
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Pikabu", "RefreshWindows", "StartStreaming", "EndStreaming", "ChangeQuality", "ChangeFps", "ChangeProperties", "MouseDown", "MouseDown2", "MouseUp", "MouseUp2", "MouseMove", "MouseMove2"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
