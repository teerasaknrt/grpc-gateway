package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"context"

	pb "github.com/teerasaknrt/grpc-gateway/api/v1"
)

type server struct {
}

func (s *server) EchoPost(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	log.Println(msg)
	log.Println(ctx)
	return msg, nil
}
func (s *server) Echo(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Println(msg, ctx)
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProfileServiceServer(grpcServer, &server{})
	log.Println("listening to port *:9090")
	grpcServer.Serve(lis)
}
