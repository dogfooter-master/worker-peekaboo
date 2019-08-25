// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "worker-peekaboo/peekaboo/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	PikabuEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.PeekabooService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{PikabuEndpoint: MakePikabuEndpoint(s)}
	for _, m := range mdw["Pikabu"] {
		eps.PikabuEndpoint = m(eps.PikabuEndpoint)
	}
	return eps
}