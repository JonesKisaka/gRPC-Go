package main

import (
	//"NAMESPACE/chat"
	"log"
	"net"

	"Repositories/gRPC-Go/chat" //import the chat file to be used here in server.go

	"google.golang.org/grpc"
)

func main() {
	//listen on port 9000 via tcp connection.(Transmission Control Protocol.)
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve  gRPC Server over port 9000: %v", err)

	}

}
