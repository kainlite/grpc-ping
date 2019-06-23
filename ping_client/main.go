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
	address     = "192.168.1.254:50000"
	defaultData = "Random data"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingServiceClient(conn)

	// Contact the server and print out its response.
	data := defaultData
	if len(os.Args) > 1 {
		// address = os.Args[1]
		data = os.Args[1]
	}

	index := 0
	for {
		trip_time := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Ping(ctx, &pb.PingRequest{Data: data})
		if err != nil {
			log.Fatalf("could not connect to: %v", err)
		}

		log.Printf("%d characters from (%s): seq=%d time=%s", len(r.Data), address, index, time.Since(trip_time))
		index++
	}
}
