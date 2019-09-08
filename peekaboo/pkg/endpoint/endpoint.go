package endpoint

import (
	"context"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"
	service "worker-peekaboo/peekaboo/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// PikabuRequest collects the request parameters for the Pikabu method.
type PikabuRequest struct {
	Req *pb.PikabuRequest `json:"req"`
}

// PikabuResponse collects the response parameters for the Pikabu method.
type PikabuResponse struct {
	Res *pb.PikabuReply `json:"res"`
	Err error           `json:"err"`
}

// MakePikabuEndpoint returns an endpoint that invokes Pikabu on the service.
func MakePikabuEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PikabuRequest)
		res, err := s.Pikabu(ctx, req.Req)
		return PikabuResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r PikabuResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Pikabu implements Service. Primarily useful in a client.
func (e Endpoints) Pikabu(ctx context.Context, req *pb.PikabuRequest) (res *pb.PikabuReply, err error) {
	request := PikabuRequest{Req: req}
	response, err := e.PikabuEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PikabuResponse).Res, response.(PikabuResponse).Err
}

// RefreshWindowsRequest collects the request parameters for the RefreshWindows method.
type RefreshWindowsRequest struct {
	Req *pb.RefreshWindowsRequest `json:"req"`
}

// RefreshWindowsResponse collects the response parameters for the RefreshWindows method.
type RefreshWindowsResponse struct {
	Res *pb.RefreshWindowsReply `json:"res"`
	Err error                   `json:"err"`
}

// MakeRefreshWindowsEndpoint returns an endpoint that invokes RefreshWindows on the service.
func MakeRefreshWindowsEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RefreshWindowsRequest)
		res, err := s.RefreshWindows(ctx, req.Req)
		return RefreshWindowsResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r RefreshWindowsResponse) Failed() error {
	return r.Err
}

// RefreshWindows implements Service. Primarily useful in a client.
func (e Endpoints) RefreshWindows(ctx context.Context, req *pb.RefreshWindowsRequest) (res *pb.RefreshWindowsReply, err error) {
	request := RefreshWindowsRequest{Req: req}
	response, err := e.RefreshWindowsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RefreshWindowsResponse).Res, response.(RefreshWindowsResponse).Err
}

// StartStreamingRequest collects the request parameters for the StartStreaming method.
type StartStreamingRequest struct {
	Req *pb.StartStreamingRequest `json:"req"`
}

// StartStreamingResponse collects the response parameters for the StartStreaming method.
type StartStreamingResponse struct {
	Res *pb.StartStreamingReply `json:"res"`
	Err error                   `json:"err"`
}

// MakeStartStreamingEndpoint returns an endpoint that invokes StartStreaming on the service.
func MakeStartStreamingEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StartStreamingRequest)
		res, err := s.StartStreaming(ctx, req.Req)
		return StartStreamingResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r StartStreamingResponse) Failed() error {
	return r.Err
}

// EndStreamingRequest collects the request parameters for the EndStreaming method.
type EndStreamingRequest struct {
	Req *pb.EndStreamingRequest `json:"req"`
}

// EndStreamingResponse collects the response parameters for the EndStreaming method.
type EndStreamingResponse struct {
	Res *pb.EndStreamingReply `json:"res"`
	Err error                 `json:"err"`
}

// MakeEndStreamingEndpoint returns an endpoint that invokes EndStreaming on the service.
func MakeEndStreamingEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EndStreamingRequest)
		res, err := s.EndStreaming(ctx, req.Req)
		return EndStreamingResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r EndStreamingResponse) Failed() error {
	return r.Err
}

// StartStreaming implements Service. Primarily useful in a client.
func (e Endpoints) StartStreaming(ctx context.Context, req *pb.StartStreamingRequest) (res *pb.StartStreamingReply, err error) {
	request := StartStreamingRequest{Req: req}
	response, err := e.StartStreamingEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(StartStreamingResponse).Res, response.(StartStreamingResponse).Err
}

// EndStreaming implements Service. Primarily useful in a client.
func (e Endpoints) EndStreaming(ctx context.Context, req *pb.EndStreamingRequest) (res *pb.EndStreamingReply, err error) {
	request := EndStreamingRequest{Req: req}
	response, err := e.EndStreamingEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(EndStreamingResponse).Res, response.(EndStreamingResponse).Err
}

// ChangeQualityRequest collects the request parameters for the ChangeQuality method.
type ChangeQualityRequest struct {
	Req *pb.ChangeQualityRequest `json:"req"`
}

// ChangeQualityResponse collects the response parameters for the ChangeQuality method.
type ChangeQualityResponse struct {
	Res *pb.ChangeQualityReply `json:"res"`
	Err error                  `json:"err"`
}

// MakeChangeQualityEndpoint returns an endpoint that invokes ChangeQuality on the service.
func MakeChangeQualityEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeQualityRequest)
		res, err := s.ChangeQuality(ctx, req.Req)
		return ChangeQualityResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r ChangeQualityResponse) Failed() error {
	return r.Err
}

// ChangeFpsRequest collects the request parameters for the ChangeFps method.
type ChangeFpsRequest struct {
	Req *pb.ChangeFpsRequest `json:"req"`
}

// ChangeFpsResponse collects the response parameters for the ChangeFps method.
type ChangeFpsResponse struct {
	Res *pb.ChangeFpsReply `json:"res"`
	Err error              `json:"err"`
}

// MakeChangeFpsEndpoint returns an endpoint that invokes ChangeFps on the service.
func MakeChangeFpsEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeFpsRequest)
		res, err := s.ChangeFps(ctx, req.Req)
		return ChangeFpsResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r ChangeFpsResponse) Failed() error {
	return r.Err
}

// ChangeQuality implements Service. Primarily useful in a client.
func (e Endpoints) ChangeQuality(ctx context.Context, req *pb.ChangeQualityRequest) (res *pb.ChangeQualityReply, err error) {
	request := ChangeQualityRequest{Req: req}
	response, err := e.ChangeQualityEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ChangeQualityResponse).Res, response.(ChangeQualityResponse).Err
}

// ChangeFps implements Service. Primarily useful in a client.
func (e Endpoints) ChangeFps(ctx context.Context, req *pb.ChangeFpsRequest) (res *pb.ChangeFpsReply, err error) {
	request := ChangeFpsRequest{Req: req}
	response, err := e.ChangeFpsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ChangeFpsResponse).Res, response.(ChangeFpsResponse).Err
}

// ChangePropertiesRequest collects the request parameters for the ChangeProperties method.
type ChangePropertiesRequest struct {
	Req *pb.ChangePropertiesRequest `json:"req"`
}

// ChangePropertiesResponse collects the response parameters for the ChangeProperties method.
type ChangePropertiesResponse struct {
	Res *pb.ChangePropertiesReply `json:"res"`
	Err error                     `json:"err"`
}

// MakeChangePropertiesEndpoint returns an endpoint that invokes ChangeProperties on the service.
func MakeChangePropertiesEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangePropertiesRequest)
		res, err := s.ChangeProperties(ctx, req.Req)
		return ChangePropertiesResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r ChangePropertiesResponse) Failed() error {
	return r.Err
}

// ChangeProperties implements Service. Primarily useful in a client.
func (e Endpoints) ChangeProperties(ctx context.Context, req *pb.ChangePropertiesRequest) (res *pb.ChangePropertiesReply, err error) {
	request := ChangePropertiesRequest{Req: req}
	response, err := e.ChangePropertiesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ChangePropertiesResponse).Res, response.(ChangePropertiesResponse).Err
}

// MouseDownRequest collects the request parameters for the MouseDown method.
type MouseDownRequest struct {
	Req *pb.MouseDownRequest `json:"req"`
}

// MouseDownResponse collects the response parameters for the MouseDown method.
type MouseDownResponse struct {
	Res *pb.MouseDownReply `json:"res"`
	Err error              `json:"err"`
}

// MakeMouseDownEndpoint returns an endpoint that invokes MouseDown on the service.
func MakeMouseDownEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MouseDownRequest)
		res, err := s.MouseDown(ctx, req.Req)
		return MouseDownResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r MouseDownResponse) Failed() error {
	return r.Err
}

// MouseDown implements Service. Primarily useful in a client.
func (e Endpoints) MouseDown(ctx context.Context, req *pb.MouseDownRequest) (res *pb.MouseDownReply, err error) {
	request := MouseDownRequest{Req: req}
	response, err := e.MouseDownEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MouseDownResponse).Res, response.(MouseDownResponse).Err
}

// MouseDown2Request collects the request parameters for the MouseDown2 method.
type MouseDown2Request struct {
	Req *pb.MouseDown2Request `json:"req"`
}

// MouseDown2Response collects the response parameters for the MouseDown2 method.
type MouseDown2Response struct {
	Res *pb.MouseDown2Reply `json:"res"`
	Err error               `json:"err"`
}

// MakeMouseDown2Endpoint returns an endpoint that invokes MouseDown2 on the service.
func MakeMouseDown2Endpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MouseDown2Request)
		res, err := s.MouseDown2(ctx, req.Req)
		return MouseDown2Response{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r MouseDown2Response) Failed() error {
	return r.Err
}

// MouseDown2 implements Service. Primarily useful in a client.
func (e Endpoints) MouseDown2(ctx context.Context, req *pb.MouseDown2Request) (res *pb.MouseDown2Reply, err error) {
	request := MouseDown2Request{Req: req}
	response, err := e.MouseDown2Endpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MouseDown2Response).Res, response.(MouseDown2Response).Err
}

// MouseUpRequest collects the request parameters for the MouseUp method.
type MouseUpRequest struct {
	Req *pb.MouseUpRequest `json:"req"`
}

// MouseUpResponse collects the response parameters for the MouseUp method.
type MouseUpResponse struct {
	Res *pb.MouseUpReply `json:"res"`
	Err error            `json:"err"`
}

// MakeMouseUpEndpoint returns an endpoint that invokes MouseUp on the service.
func MakeMouseUpEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MouseUpRequest)
		res, err := s.MouseUp(ctx, req.Req)
		return MouseUpResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r MouseUpResponse) Failed() error {
	return r.Err
}

// MouseUp2Request collects the request parameters for the MouseUp2 method.
type MouseUp2Request struct {
	Req *pb.MouseUp2Request `json:"req"`
}

// MouseUp2Response collects the response parameters for the MouseUp2 method.
type MouseUp2Response struct {
	Res *pb.MouseUp2Reply `json:"res"`
	Err error             `json:"err"`
}

// MakeMouseUp2Endpoint returns an endpoint that invokes MouseUp2 on the service.
func MakeMouseUp2Endpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MouseUp2Request)
		res, err := s.MouseUp2(ctx, req.Req)
		return MouseUp2Response{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r MouseUp2Response) Failed() error {
	return r.Err
}

// MouseMoveRequest collects the request parameters for the MouseMove method.
type MouseMoveRequest struct {
	Req *pb.MouseMoveRequest `json:"req"`
}

// MouseMoveResponse collects the response parameters for the MouseMove method.
type MouseMoveResponse struct {
	Res *pb.MouseMoveReply `json:"res"`
	Err error              `json:"err"`
}

// MakeMouseMoveEndpoint returns an endpoint that invokes MouseMove on the service.
func MakeMouseMoveEndpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MouseMoveRequest)
		res, err := s.MouseMove(ctx, req.Req)
		return MouseMoveResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r MouseMoveResponse) Failed() error {
	return r.Err
}

// MouseMove2Request collects the request parameters for the MouseMove2 method.
type MouseMove2Request struct {
	Req *pb.MouseMove2Request `json:"req"`
}

// MouseMove2Response collects the response parameters for the MouseMove2 method.
type MouseMove2Response struct {
	Res *pb.MouseMove2Reply `json:"res"`
	Err error               `json:"err"`
}

// MakeMouseMove2Endpoint returns an endpoint that invokes MouseMove2 on the service.
func MakeMouseMove2Endpoint(s service.PeekabooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MouseMove2Request)
		res, err := s.MouseMove2(ctx, req.Req)
		return MouseMove2Response{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r MouseMove2Response) Failed() error {
	return r.Err
}

// MouseUp implements Service. Primarily useful in a client.
func (e Endpoints) MouseUp(ctx context.Context, req *pb.MouseUpRequest) (res *pb.MouseUpReply, err error) {
	request := MouseUpRequest{Req: req}
	response, err := e.MouseUpEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MouseUpResponse).Res, response.(MouseUpResponse).Err
}

// MouseUp2 implements Service. Primarily useful in a client.
func (e Endpoints) MouseUp2(ctx context.Context, req *pb.MouseUp2Request) (res *pb.MouseUp2Reply, err error) {
	request := MouseUp2Request{Req: req}
	response, err := e.MouseUp2Endpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MouseUp2Response).Res, response.(MouseUp2Response).Err
}

// MouseMove implements Service. Primarily useful in a client.
func (e Endpoints) MouseMove(ctx context.Context, req *pb.MouseMoveRequest) (res *pb.MouseMoveReply, err error) {
	request := MouseMoveRequest{Req: req}
	response, err := e.MouseMoveEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MouseMoveResponse).Res, response.(MouseMoveResponse).Err
}

// MouseMove2 implements Service. Primarily useful in a client.
func (e Endpoints) MouseMove2(ctx context.Context, req *pb.MouseMove2Request) (res *pb.MouseMove2Reply, err error) {
	request := MouseMove2Request{Req: req}
	response, err := e.MouseMove2Endpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(MouseMove2Response).Res, response.(MouseMove2Response).Err
}
