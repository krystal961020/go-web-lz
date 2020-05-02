package main

//import (
//	"github.com/go-kit/kit/endpoint"
//	"github.com/go-kit/kit/sd"
//	grpctransport "github.com/go-kit/kit/transport/grpc"
//	"google.golang.org/grpc"
//	"io"
//	ggrpc "kit/test/grpc"
//	"kit/test/grpc/pb"
//)
//
//type CalculateRequest struct {
//	RequestType string `json:"request_type"`
//	A           int    `json:"a"`
//	B           int    `json:"b"`
//}
//
//// ArithmeticResponse define response struct
//type CalculateResponse struct {
//	Result int   `json:"result"`
//	Error  error `json:"error"`
//}
//
//
//func CalculateFactory( options []grpctransport.ClientOption) sd.Factory {
//	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
//		conn, err := grpc.Dial(instance, grpc.WithInsecure())
//		if err != nil {
//			return nil, nil, err
//		}
//		var (
//			enc grpctransport.EncodeRequestFunc
//			dec grpctransport.DecodeResponseFunc
//		)
//		enc, dec = ggrpc.EncodeGRPCCalculateRequest, ggrpc.DecodeGRPCCalculateResponse
//		return grpctransport.NewClient(conn, "test","calculate", enc, dec,pb.CalculateResponse{},options...).Endpoint(), nil, nil
//	}
//
//	}
//
