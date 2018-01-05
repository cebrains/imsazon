package mail

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

type sendRequest struct {
	SmtpTemplate			SmtpTemplateData
}

type sendResponse struct {
	Err error				`json:"error,omitempty"`
}

func(r sendResponse) error() error { return r.Err }

func makeSendEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(sendRequest)
		err := s.Send(req.SmtpTemplate)
		return sendResponse{Err: err}, nil
	}
}