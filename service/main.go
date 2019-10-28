package main

import (
	"errors"
	"net"

	"go.uber.org/zap"

	context "context"

	pb "github.com/teerasaknrt/grpc-gateway/api/v1"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	zapLogger *zap.Logger
)

type server struct {
}

func (s *server) EchoPost(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	//log.Println(msg)
	if msg.Value == "" {
		err := errors.New("serverUnavailable")
		return nil, err
	}
	return msg, nil
}
func (s *server) Echo(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	//log.Println(msg, ctx)
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := newGRPCServer()
	pb.RegisterProfileServiceServer(grpcServer, &server{})
	log.Println("listening to port *:9090")
	grpcServer.Serve(lis)
}

func newGRPCServer() *grpc.Server {
	logrusEntry := log.NewEntry(log.StandardLogger())
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		)),
	)

	reflection.Register(server)

	grpc_prometheus.Register(server)

	return server
}
