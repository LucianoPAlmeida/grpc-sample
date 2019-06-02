package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "drivers.org"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) FindDriver(ctx context.Context, in *pb.DriverRequest) (*pb.DriverResponse, error) {
	log.Printf("Received: %d", in.Number)

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDriverServer(s, &server{})
	log.Printf("gRPC server starting on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
