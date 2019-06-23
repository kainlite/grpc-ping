package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/kainlite/grpc-ping/ping"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50000"
	defaultData = "Random data"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingClient(conn)

	// Contact the server and print out its response.
	data := defaultName
	if len(os.Args) > 2 {
		address = os.Args[1]
		data = os.Args[2]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{Data: data})
	if err != nil {
		log.Fatalf("could not connect to: %v", err)
	}
	log.Printf("Response: %s", r.Message)
}
