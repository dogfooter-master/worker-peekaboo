package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"
	endpoint "worker-peekaboo/peekaboo/pkg/endpoint"
	"worker-peekaboo/peekaboo/pkg/service"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

func WebSocketFunc(w http1.ResponseWriter, r *http1.Request) {
	service.ServeWebSocket(w, r)
}

// makePikabuHandler creates the handler logic
func makePikabuHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/pikabu").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PikabuEndpoint, decodePikabuRequest, encodePikabuResponse, options...)))
	m.HandleFunc("/ws", WebSocketFunc)
}

// decodePikabuRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePikabuRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PikabuRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePikabuResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePikabuResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeRefreshWindowsHandler creates the handler logic
func makeRefreshWindowsHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/refresh-windows").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.RefreshWindowsEndpoint, decodeRefreshWindowsRequest, encodeRefreshWindowsResponse, options...)))
}

// decodeRefreshWindowsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRefreshWindowsRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RefreshWindowsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRefreshWindowsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRefreshWindowsResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeStartStreamingHandler creates the handler logic
func makeStartStreamingHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/start-streaming").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.StartStreamingEndpoint, decodeStartStreamingRequest, encodeStartStreamingResponse, options...)))
}

// decodeStartStreamingRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeStartStreamingRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.StartStreamingRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeStartStreamingResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeStartStreamingResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeEndStreamingHandler creates the handler logic
func makeEndStreamingHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/end-streaming").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.EndStreamingEndpoint, decodeEndStreamingRequest, encodeEndStreamingResponse, options...)))
}

// decodeEndStreamingRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeEndStreamingRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.EndStreamingRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeEndStreamingResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeEndStreamingResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeChangeQualityHandler creates the handler logic
func makeChangeQualityHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/change-quality").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ChangeQualityEndpoint, decodeChangeQualityRequest, encodeChangeQualityResponse, options...)))
}

// decodeChangeQualityRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeChangeQualityRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ChangeQualityRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeChangeQualityResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeChangeQualityResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeChangeFpsHandler creates the handler logic
func makeChangeFpsHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/change-fps").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ChangeFpsEndpoint, decodeChangeFpsRequest, encodeChangeFpsResponse, options...)))
}

// decodeChangeFpsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeChangeFpsRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ChangeFpsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeChangeFpsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeChangeFpsResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeChangePropertiesHandler creates the handler logic
func makeChangePropertiesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/change-properties").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ChangePropertiesEndpoint, decodeChangePropertiesRequest, encodeChangePropertiesResponse, options...)))
}

// decodeChangePropertiesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeChangePropertiesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ChangePropertiesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeChangePropertiesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeChangePropertiesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeMouseDownHandler creates the handler logic
func makeMouseDownHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/mouse-down").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.MouseDownEndpoint, decodeMouseDownRequest, encodeMouseDownResponse, options...)))
}

// decodeMouseDownRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMouseDownRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.MouseDownRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMouseDownResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMouseDownResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeMouseDown2Handler creates the handler logic
func makeMouseDown2Handler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/mouse-down2").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.MouseDown2Endpoint, decodeMouseDown2Request, encodeMouseDown2Response, options...)))
}

// decodeMouseDown2Request is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMouseDown2Request(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.MouseDown2Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMouseDown2Response is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMouseDown2Response(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeMouseUpHandler creates the handler logic
func makeMouseUpHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/mouse-up").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.MouseUpEndpoint, decodeMouseUpRequest, encodeMouseUpResponse, options...)))
}

// decodeMouseUpRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMouseUpRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.MouseUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMouseUpResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMouseUpResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeMouseUp2Handler creates the handler logic
func makeMouseUp2Handler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/mouse-up2").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.MouseUp2Endpoint, decodeMouseUp2Request, encodeMouseUp2Response, options...)))
}

// decodeMouseUp2Request is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMouseUp2Request(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.MouseUp2Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMouseUp2Response is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMouseUp2Response(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeMouseMoveHandler creates the handler logic
func makeMouseMoveHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/mouse-move").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.MouseMoveEndpoint, decodeMouseMoveRequest, encodeMouseMoveResponse, options...)))
}

// decodeMouseMoveRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMouseMoveRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.MouseMoveRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMouseMoveResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMouseMoveResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeMouseMove2Handler creates the handler logic
func makeMouseMove2Handler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/mouse-move2").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.MouseMove2Endpoint, decodeMouseMove2Request, encodeMouseMove2Response, options...)))
}

// decodeMouseMove2Request is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeMouseMove2Request(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.MouseMove2Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeMouseMove2Response is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeMouseMove2Response(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
