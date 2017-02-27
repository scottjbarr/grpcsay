package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/scottjbarr/grpcsay"
	pb "github.com/scottjbarr/grpcsay/grpcsay"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Say(ctx context.Context,
	in *pb.SayRequest) (*pb.SayResponse, error) {

	out := grpcsay.Say(in.Message)

	return &pb.SayResponse{Message: out}, nil
}

func main() {
	bind := os.Getenv("BIND")

	if len(bind) == 0 {
		fmt.Printf("Usage: BIND=:50051 bovine_server")
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGopherServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
