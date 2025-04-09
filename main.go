package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/demo-grpc/demo"
	"google.golang.org/grpc"
)

type myDemoServer struct {
	demo.UnimplementedDemoServer
}

func (s myDemoServer) Create(ctx context.Context, r *demo.CreateRequest) (*demo.CreateResponse, error) {
	fmt.Println("Create-Requested", r.Message)
	method, ok := grpc.Method(ctx)
	if ok {
		// method will be something like "/package.Service/Method"
		fmt.Println("Requested method:", method)
	}

	return &demo.CreateResponse{Message: r.Message}, nil
}

func (s myDemoServer) Check(_ context.Context, r *demo.CreateRequest) (*demo.CreateResponse, error) {
	fmt.Println("Check-Requested", r.Message)
	return &demo.CreateResponse{Message: r.Message}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Cannot create listner %s", err.Error())
		return
	}

	server := grpc.NewServer()
	service := &myDemoServer{}
	demo.RegisterDemoServer(server, service)

	fmt.Println("Listening on port:", 9000)
	err = server.Serve(listner)
	if err != nil {
		log.Fatalf("Cannot create listner %s", err.Error())
		return
	}

	fmt.Println("Listening on port 9000")
}
