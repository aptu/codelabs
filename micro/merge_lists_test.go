package main

import (
    "testing"
	"context"
	"log"
	"micro/pb"

	grpc "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)


func TestServer(t *testing.T) {

    // conn, err := grpc.Dial("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
    conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
    defer conn.Close()

    if err != nil {
        log.Println(err)
    }

    client := pb.NewMergeListsClient(conn)
    res, err := client.Merge(context.Background(), &pb.MergeListRequest{
        List1: &pb.List{V:[]int32 {1, 3, 5, 7}},
        List2: &pb.List{V:[]int32 {2, 4, 6, 8}},
        },
    )

    log.Println(res)
    log.Println(err)
}