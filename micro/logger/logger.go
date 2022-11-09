package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"logger/pb"
	"net"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedLoggerServer
}

func newServer() *server {
	s := &server{}

	return s
}

func (s *server) Info(ctx context.Context, in *pb.LogRequest) (*empty.Empty, error) {
	fmt.Println("[INFO] ", in.Text)
	return &empty.Empty{}, nil
}

func (s *server) Error(ctx context.Context, in *pb.LogRequest) (*empty.Empty, error) {
	fmt.Println("[ERROR] ", in.Text)
	return &empty.Empty{}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 7778))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    10 * time.Second,
		Timeout: 20 * time.Second,
	}))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLoggerServer(grpcServer, newServer())
	reflection.Register(grpcServer)
	log.Println("Starting server...")
	grpcServer.Serve(lis)
}
