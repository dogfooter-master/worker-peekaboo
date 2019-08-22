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
