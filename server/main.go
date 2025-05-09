package main

import (
	"context"
	"log"
	"net"

	pb "github.com/iuketaylor/grpc_programmes/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedProgrammeServiceServer
}

func (s *server) GetProgramme(ctx context.Context, req *pb.GetProgrammeRequest) (*pb.Programme, error) {
	log.Printf("Received request for PID %s", req.Pid)
	var mockDatabase = map[string]*pb.Programme{
		"b006m86d": {
			Id:    "b006m86d",
			Live:  false,
			Type:  "programme",
			Title: "Eastenders",
			Images: &pb.Image{
				Type:     "image",
				Standard: "https://ichef.bbci.co.uk/images/ic/640x360/p018hq72.jpg",
			},
		},
		"p1234567": {
			Id:    "p1234567",
			Live:  true,
			Type:  "programme",
			Title: "Planet Earth",
			Images: &pb.Image{
				Type:     "image",
				Standard: "https://ichef.bbci.co.uk/images/ic/640x360/p0987654.jpg",
			},
		},
	}

	programme, ok := mockDatabase[req.Pid]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Programme not found %s", req.Pid)
	}

	return programme, nil
}

func main() {

	s := grpc.NewServer()
	pb.RegisterProgrammeServiceServer(s, &server{})

	reflection.Register(s)

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	log.Println("gRPC server listening on port :5001")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
