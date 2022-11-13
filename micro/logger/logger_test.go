package main

import (
	"context"
	"log"
	"logger/pb"
	"testing"

	grpc "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

func TestServer(t *testing.T) {

	conn, err := grpc.Dial("0.0.0.0:7778", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	client := pb.NewLoggerClient(conn)
	res, err := client.Info(context.Background(), &pb.LogRequest{
		Text: "Testing"})

	log.Println(res)
	log.Println(err)
}
