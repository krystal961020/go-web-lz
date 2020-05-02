package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	lb "kit/test/grpc/consul"
	"kit/test/grpc/pb"
	"log"
	"time"
)

func main() {
	var (
		_ = flag.Int("n", 0, "Number of calls to service")
		t = flag.Duration("t", 1*time.Second, "Sleep interval between calls")
	)
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Lookup service in Consul
	cli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	services, err := cli.Agent().Services()
	service := services["Calculate-127.0.0.1-9002"]
	log.Printf("%v", service)

	// Resolver for the "echo" service
	r, err := lb.NewResolver(cli, service.Service, "")
	if err != nil {
		log.Fatal(err)
	}

	// Dial options
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// Enabling WithBlock tells the client to not give up trying to find a server
	opts = append(opts, grpc.WithBlock())
	// However, we're still setting a timeout so that if the server takes too long, we still give up
	opts = append(opts, grpc.WithTimeout(10*time.Second))
	// Add resolver with RoundRobin balancer here
	opts = append(opts, grpc.WithBalancer(grpc.RoundRobin(r)))

	// Notice the blank address
	conn, err := grpc.Dial("", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewCalculateServiceClient(conn)

	timeout := time.Duration(int64(3+(*t).Seconds())) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	req := &pb.CalculateRequest{
		RequestType: "Add",
		A:           2,
		B:           3,
	}
	res, err := client.Calculate(ctx, req)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	fmt.Printf("%s%v\n", "krystal :", res)
	time.Sleep(*t)
	cancel()
}

/*
func main1() {
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
*/
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
