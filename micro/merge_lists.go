package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"micro/pb"
	"net"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMergeListsServer
}

func newServer() *server {
	s := &server{}

	return s
}

func (s *server) Merge(ctx context.Context, request *pb.MergeListRequest) (*pb.MergeListResponse, error) {
	log.Println("Merging lists...")
	log.Println(request)
	start := time.Now()
	res := merge(request.List1.GetV(), request.List2.GetV())
	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Runtime: ", elapsed)
	//return &pb.MergeListResponse{Merged: &pb.List{V: []int32{3}}}, nil
	return &pb.MergeListResponse{Merged: &pb.List{V: res}}, nil
}

func merge(l1 []int32, l2 []int32) []int32 {
	l3 := make([]int32, len(l1)+len(l2))
	r1 := 0
	r2 := 0
	w := 0

	for r1 < len(l1) && r2 < len(l2) {
		if l1[r1] < l2[r2] {
			l3[w] = l1[r1]
			r1++
		} else {
			l3[w] = l2[r2]
			r2++
		}
		w++
	}

	if r1 >= len(l1) {
		copy(l3[w:], l2[r2:])
	}
	if r2 >= len(l2) {
		copy(l3[w:], l1[r1:])
	}

	return l3
}

func main() {
	// l1 := []int{1, 3, 5, 6}
	// l2 := []int{2, 4, 6, 8}
	// fmt.Println(l1)
	// fmt.Println(l2)
	// l3 := merge(l1, l2)
	// fmt.Println(l3)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    10 * time.Second,
		Timeout: 20 * time.Second,
	}))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMergeListsServer(grpcServer, newServer())
	reflection.Register(grpcServer)
	log.Println("Starting server...")
	grpcServer.Serve(lis)
}
