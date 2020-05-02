package main

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"io"
	ggrpc "kit/test/grpc"
	"time"
)

func New(consulAddr string, logger log.Logger) (ggrpc.Service, error) {
	var (
		consulTags   = []string{"grpc"}
		passingOnly  = true
		retryMax     = 3
		retryTimeout = 500 * time.Millisecond
	)

	conClient, err := api.NewClient(&api.Config{
		Address: consulAddr,
	})
	if err != nil {
		return nil, err
	}
	sdClient := consul.NewClient(conClient)
	var (
		registerSrvName = "Calculate"
		instanter       = consul.NewInstancer(sdClient, logger, registerSrvName, consulTags, passingOnly)
		endpoints       = ggrpc.CalculateEndpoints{}
	)
	{
		factory := factoryFor(ggrpc.MakeCalculateEndpoint)
		endpointer := sd.NewEndpointer(instanter, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CalculateEndpoint = retry

	}
	return endpoints, nil
}

func factoryFor(makeEndpoint func(ggrpc.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
		if err != nil {
			return nil, nil, err
		}
		//ctx := context.Background()
		//srv1 := pb.NewCalculateServiceClient(conn)
		//
		//// Contact the server and print out its response.
		////name := defaultName
		//result, err := srv1.Calculate(ctx, &pb.CalculateRequest{RequestType:"Add",A:1,B:9})
		//if err != nil {
		//	fmt.Println("could not greet: %v", err)
		//}
		//fmt.Println("Greeting: %s", result)
		srv, err := newGRPCClient(conn, []grpctransport.ClientOption{})
		if err != nil {
			return nil, nil, err
		}
		endpoints := makeEndpoint(srv)
		return endpoints, conn, err
	}
}

func newGRPCClient(conn *grpc.ClientConn, options []grpctransport.ClientOption) (ggrpc.Service, error) {
	var CalculateEndpoint endpoint.Endpoint
	{
		CalculateEndpoint = grpctransport.NewClient(conn, "Calculate", "Calculate", ggrpc.EncodeGRPCCalculateRequest, ggrpc.DecodeGRPCCalculateResponse, ggrpc.CalculateResponse{}, options...).Endpoint()
	}

	return ggrpc.CalculateEndpoints{CalculateEndpoint: CalculateEndpoint}, nil
}

//func EncodeGRPCCalculateRequest(_ context.Context, r interface{}) (interface{}, error) {
//	req := r.(CalculateRequest)
//	return &pb.CalculateRequest{
//		RequestType: req.RequestType,
//		A:           int32(req.A),
//		B:           int32(req.B),
//	}, nil
//}
//
//func DecodeGRPCCalculateResponse(_ context.Context, r interface{}) (interface{}, error) {
//	resp := r.(*pb.CalculateResponse)
//	return CalculateResponse{
//		Result: int(resp.Result),
//		Error:  errors.New(resp.Err),
//	}, nil
//}
