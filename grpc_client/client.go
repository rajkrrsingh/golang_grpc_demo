package main

import (
	"context"
	"fmt"
	"github.com/rajkrrsingh/golang_grpc_demo/grpcpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	opts := grpc.WithInsecure()

	cc,err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := grpcpb.NewServerInfoClient(cc)

	callGetServerInfo(c)
}

func callGetServerInfo(client grpcpb.ServerInfoClient) {
	fmt.Println("[callGetServerInfo]  client is about to send GetServerInfo")

	req := &grpcpb.GetServerInfoRequest{}

	res,err := client.GetServerInfo(context.Background(),req)
	if err != nil {
		log.Fatalf("error while calling GetServerInfo RPC: %v", err)
	}
	fmt.Printf("[callGetServerInfo] Response : hostname : %s current server time : %s", res.Hostname, res.Currentservertime)

}
