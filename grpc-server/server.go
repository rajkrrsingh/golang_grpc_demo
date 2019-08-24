package main

import (
	"context"
	"fmt"
	"github.com/rajkrrsingh/golang_grpc_demo/grpcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"time"
)

type server struct {

}

func (*server) GetServerInfo(ctx context.Context, req *grpcpb.GetServerInfoRequest) (*grpcpb.GetServerResponse, error) {
	fmt.Println("[GetServerInfo] serving GetServerInfo request")
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Unable to get the hostname %v",err)
		return nil,status.Error(codes.Internal,"Unable to get the hostname")
	}
	res := &grpcpb.GetServerResponse{
		Hostname:             hostname,
		Currentservertime:    time.Now().String(),
	}
	return res,nil
}


// go generate endpoint: rajkumar.singh@HW14317 golang_grpc_demo $ ~/Downloads/protoc-3.9.1-osx-x86_64/bin/protoc grpcpb/app.proto --go_out=plugins=grpc:.
func main() {
	fmt.Println("GRPC Server starting!!")
	list, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatal("unable to listen on port %v", err)
	}

	s := grpc.NewServer()
	grpcpb.RegisterServerInfoServer(s,&server{})

	error := s.Serve(list)
	if error != nil {
		log.Fatal("failed to server %v", error)
	}
}
