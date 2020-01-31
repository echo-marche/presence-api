package main

import (
	"log"
	"net"

	pb "github.com/alduh-co/presence-api/proto/pb"
	"github.com/alduh-co/presence-api/servers"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServer(server, &servers.UserServer{})
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
