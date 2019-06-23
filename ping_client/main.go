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
	defaultAddress = "localhost:50000"
	defaultData    = "00"
)

func main() {
	data := defaultData
	address := defaultAddress
	if len(os.Args) > 2 {
		address = os.Args[1]
		data = os.Args[2]
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingServiceClient(conn)

	index := 0
	for {
		trip_time := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Ping(ctx, &pb.PingRequest{Data: data})
		if err != nil {
			log.Fatalf("could not connect to: %v", err)
		}

		log.Printf("%d characters roundtrip to (%s): seq=%d time=%s", len(r.Data), address, index, time.Since(trip_time))
		time.Sleep(1 * time.Second)
		index++
	}
}
