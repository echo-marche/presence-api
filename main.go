package main

import (
	"log"
	"net"

	"github.com/echo-marche/presence-api/config"
	"github.com/echo-marche/presence-api/infrastructure"
	pb "github.com/echo-marche/presence-api/proto/pb"
	"github.com/echo-marche/presence-api/servers"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// DB connection
	dbMap := infrastructure.NewDbMap()
	defer dbMap.Db.Close()

	// init gRPC server
	listenPort, err := net.Listen("tcp", ":"+config.GetEnv("PRESENCE_API_SERVICE_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
		)),
	)
	presenceServer := &servers.PresenceServer{DbMap: dbMap}
	healthServer := &servers.HealthServer{}
	pb.RegisterPresenceServer(server, presenceServer)
	health.RegisterHealthServer(server, healthServer)
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
