// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
	endpoint "worker-peekaboo/peekaboo/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makePikabuHandler(m, endpoints, options["Pikabu"])
	makeRefreshWindowsHandler(m, endpoints, options["RefreshWindows"])
	makeStartStreamingHandler(m, endpoints, options["StartStreaming"])
	makeEndStreamingHandler(m, endpoints, options["EndStreaming"])
	makeChangeQualityHandler(m, endpoints, options["ChangeQuality"])
	makeChangeFpsHandler(m, endpoints, options["ChangeFps"])
	makeChangePropertiesHandler(m, endpoints, options["ChangeProperties"])
	return m
}
