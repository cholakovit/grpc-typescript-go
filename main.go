package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"


	pb "gots/proto/client"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8082, "The server port")
)

type server struct {
	pb.UnimplementedRandomServer
}

// PingPong implements the PingPong method of the client.RandomServer interface.
func (s *server) PingPong(ctx context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	log.Printf("Received: %v\n", in.Message)
	
	// Implement your logic here to generate the PongResponse.
	response := &pb.PongResponse{Message: "Pong"}

	return response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	//pb.RegisterRandomServer(s, &server{})
	pb.RegisterRandomServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}