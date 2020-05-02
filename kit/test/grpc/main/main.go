package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	. "kit/test/grpc"
	"kit/test/grpc/pb"
	//"github.com/micro/util/go/lib/net"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc Service
	svc = CalculateService{}

	// add logging middleware to service

	calEndpoint := MakeCalculateEndpoint(svc)
	endpts := CalculateEndpoints{
		CalculateEndpoint: calEndpoint,
	}

	consulAddres := fmt.Sprintf("%s:%s", "http://127.0.0.1", "8500")
	//grpcPort, _ := strconv.Atoi("9002")
	//metricsPort, _ := strconv.Atoi("8000")
	consulReg := NewConsulRegister(consulAddres, "Calculate", "127.0.0.1", 9002, []string{"grpc"})
	register, err := consulReg.NewConsulGRPCRegister()
	defer register.Deregister()

	if err != nil {
		level.Error(logger).Log(
			"consulAddres", consulAddres,
			//"serviceName", cfg.serviceName,
			//"grpcPort", grpcPort,
			//"metricsPort", metricsPort,
			//"tags", []string{cfg.nameSpace, cfg.serviceName},
			"err", err,
		)
	}
	//register.Register()

	//grpc server
	go func() {
		listener, err := net.Listen("tcp", ":9002")
		if err != nil {
			errChan <- err
			return
		}
		handler := NewGRPCServer(endpts)
		gRPCServer := grpc.NewServer()
		pb.RegisterCalculateServiceServer(gRPCServer, handler)
		grpc_health_v1.RegisterHealthServer(gRPCServer, &HealthImpl{})
		register.Register()
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	fmt.Println(error)
}
