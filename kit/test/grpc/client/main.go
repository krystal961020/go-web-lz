package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"kit/test/grpc/pb"
	"time"
)

func main1() {
	var (
		grpcAddr = flag.String("addr", "127.0.0.1:9002", "gRPC address")
	)
	flag.Parse()
	ctx := context.Background()

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		fmt.Println("gRPC dial err:", err)
	}
	defer conn.Close()

	svr := pb.NewCalculateServiceClient(conn)
	result, err := svr.Calculate(ctx, &pb.CalculateRequest{RequestType: "Add", A: 1, B: 9})
	if err != nil {
		fmt.Println("calculate error", err.Error())

	}

	fmt.Println("result=", result)
}

func main() {
	//main1()
	//var logger log.Logger
	//logFormat := log.LogfmtLogger(os.Stdout)
	//logger = log.LoggerFunc(func (keyvals ...interface{}) error {
	//	if err := logFormat.Log(keyvals...); err != nil {
	//		// handle error
	//	}
	//	return nil
	//})
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//
	//defer cancel()
	s, err := New("127.0.0.1:8500", log.NewNopLogger())
	if err != nil {
		fmt.Println(err)
		//logger.Log("service client error: ", err.Error())
	}

	res, err1 := s.Calculate(ctx, "Add", 1, 2)
	if err1 != nil {
		fmt.Println(res)
	} else {
		fmt.Println(err1)
	}
	fmt.Println(res)
}

//import (
//	"context"
//	"google.golang.org/grpc"
//	"kit/test/grpc/consul"
//	pb "kit/test/grpc/pb"
//	"log"
//)
//
//const (
//	target      = "http://127.0.0.1:8500"
//	defaultName = "tttt"
//)
//
//func main() {
//	consul.Init()
//	// Set up a connection to the server.
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithDefaultServiceConfig("round_robin"))
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	srv := pb.NewCalculateServiceClient(conn)
//
//	// Contact the server and print out its response.
//	//name := defaultName
//	result, err := srv.Calculate(ctx, &pb.CalculateRequest{RequestType:"Add",A:1,B:9})
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	log.Printf("Greeting: %s", result)
//	//for {
//	//	ctx, _ := context.WithTimeout(context.Background(), time.Second)
//	//	result, err := srv.Calculate(ctx, &pb.CalculateRequest{RequestType:"Add",A:1,B:9})
//	//	if err != nil {
//	//		log.Fatalf("could not greet: %v", err)
//	//	}
//	//	log.Printf("Greeting: %s", result)
//	//	time.Sleep(time.Second * 2)
//	//}
//}
