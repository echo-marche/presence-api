package main

import (
	"log"
	"net"

	"github.com/alduh-co/presence-api/config"
	pb "github.com/alduh-co/presence-api/proto/pb"
	"github.com/alduh-co/presence-api/servers"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":"+config.GetEnv("API_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	pb.RegisterPresenceServer(server, &servers.PresenceServer{})
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
