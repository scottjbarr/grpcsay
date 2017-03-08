package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/scottjbarr/grpcsay/grpcsay"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func usage() string {
	return "Usage: ADDRESS=localhost:50051 grpcsay_client \"Message to send\""
}
func main() {
	address := os.Getenv("ADDRESS")

	if len(address) == 0 {
		fmt.Printf("%v\n", usage())
		os.Exit(1)
	}

	// get the message to send from args
	message := os.Args[1]
	if len(message) == 0 {
		fmt.Printf(usage())
		os.Exit(2)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGopherClient(conn)

	// Contact the server and print out its response.
	req := pb.SayRequest{Message: message}
	r, err := c.Say(context.Background(), &req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("%s\n", r.Message)
}
