package main

import (
	"log"
	"net"

	"github.com/alduh-co/presence-api/config"
	"github.com/alduh-co/presence-api/infrastructure"
	pb "github.com/alduh-co/presence-api/proto/pb"
	"github.com/alduh-co/presence-api/servers"
	"google.golang.org/grpc"
)

func main() {
	// DB connection
	db := infrastructure.NewDb()
	defer db.Close()

	// init gRPC server
	listenPort, err := net.Listen("tcp", ":"+config.GetEnv("API_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	presenceServer := &servers.PresenceServer{Db: db}
	pb.RegisterPresenceServer(server, presenceServer)
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
