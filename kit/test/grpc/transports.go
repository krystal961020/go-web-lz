package grpc

import (
	"context"
	"errors"
	_ "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"kit/test/grpc/pb"
)

type grpcServer struct {
	calculate grpc.Handler
}

func (s *grpcServer) Calculate(ctx context.Context, r *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	_, resp, err := s.calculate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CalculateResponse), nil
}

func NewGRPCServer(endpoints CalculateEndpoints) pb.CalculateServiceServer {
	return &grpcServer{
		calculate: grpc.NewServer(
			endpoints.CalculateEndpoint,
			DecodeGRPCCalculateRequest,
			EncodeGRPCCalculateResponse,
		),
	}
}

func DecodeGRPCCalculateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CalculateRequest)
	return CalculateRequest{
		RequestType: req.RequestType,
		A:           int(req.A),
		B:           int(req.B),
	}, nil
}

func EncodeGRPCCalculateResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(CalculateResponse)

	if resp.Error != nil {
		return &pb.CalculateResponse{
			Result: int32(resp.Result),
			Err:    resp.Error.Error(),
		}, nil
	}

	return &pb.CalculateResponse{
		Result: int32(resp.Result),
		Err:    "",
	}, nil
}

func EncodeGRPCCalculateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(CalculateRequest)
	return &pb.CalculateRequest{
		RequestType: req.RequestType,
		A:           int32(req.A),
		B:           int32(req.B),
	}, nil
}

func DecodeGRPCCalculateResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.CalculateResponse)
	return CalculateResponse{
		Result: int(resp.Result),
		Error:  errors.New(resp.Err),
	}, nil
}
