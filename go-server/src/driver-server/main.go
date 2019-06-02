package main

import (
	"context"
	"log"
	"net"

	pb "drivers.org"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) FindDriver(ctx context.Context, in *pb.DriverRequest) (*pb.DriverResponse, error) {
	log.Printf("Received: %d", in.Number)
	drivers := map[int32]string{
		44: "Lewis Hamilton",
		77: "Valtteri Bottas",
		5:  "Sebastian Vettel",
		16: "Charles Leclerc",
		10: "Pierre Gasly",
		33: "Max Verstappen",
		3:  "Daniel Ricciardo",
		27: "Nico Hulkenberg",
		8:  "Romain Grosjean",
		20: "Kevin Magnussen",
		55: "Carlos Sainz",
		4:  "Lando Norris",
		11: "Sergio Perez",
		18: "Lance Stroll",
		7:  "Kimi Raikkonen",
		99: "Antonio Giovinazzi",
		26: "Daniil Kvyat",
		23: "Alexander Albon",
		63: "George Russell",
		88: "Robert Kubica",
	}
	name := drivers[in.Number]
	if name == "" {
		return nil, nil
	}
	return &pb.DriverResponse{Name: name}, nil
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
