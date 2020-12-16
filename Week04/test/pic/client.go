package main

import (
	"context"
	"log"
	"time"

	pb "Week04/api/pic/v1"

	"google.golang.org/grpc"
)

type MockReq struct {
	addr string
	id int32
}

var mkReq *MockReq = &MockReq{":8080", 1}

func main() {
	conn, err := grpc.Dial(mkReq.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	defer conn.Close()

	picClient := pb.NewPicClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := picClient.GetPicInfoById(ctx, &pb.GetPicInfoByIdRequest{Id: mkReq.id})
	if err != nil {
		log.Fatalf("Pic request register failed: %v", err)
	}
	log.Printf("Pic request id: %d, got name: %s, url: %s", resp.GetId(), resp.GetName(), resp.GetUrl())
}