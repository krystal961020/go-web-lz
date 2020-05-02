package grpc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// CalculateEndpoint define endpoint
type CalculateEndpoints struct {
	CalculateEndpoint   endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

// ArithmeticRequest define request struct
type CalculateRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

// ArithmeticResponse define response struct
type CalculateResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

// MakeArithmeticEndpoint make endpoint
func MakeCalculateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CalculateRequest)

		var (
			res, a, b int
			calError  error
		)

		a = req.A
		b = req.B

		res, calError = svc.Calculate(ctx, req.RequestType, a, b)

		return CalculateResponse{Result: res, Error: calError}, nil
	}
}

func (ae CalculateEndpoints) Calculate(ctx context.Context, reqType string, a, b int) (res int, err error) {
	//ctx = context.Background()

	resp, err := ae.CalculateEndpoint(ctx, CalculateRequest{
		RequestType: reqType,
		A:           a,
		B:           b,
	})
	if err != nil {
		return 0, err
	}
	response := resp.(CalculateResponse)
	return response.Result, nil
}

//func (ae CalculateEndpoints) HealthCheck() bool {
//	return false
//}
