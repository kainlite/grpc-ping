package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kainlite/grpc-ping/ping"
	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

// server is used to implement ping.PingServer.
type server struct{}

// Ping implements ping.PingServer
func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", in.Data)
	return &pb.PingResponse{Data: "Data: " + in.Data}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
