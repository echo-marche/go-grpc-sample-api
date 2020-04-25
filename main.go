package main

import (
	"log"
	"net"

	"github.com/echo-marche/hack-tech-tips-api/config"
	"github.com/echo-marche/hack-tech-tips-api/infrastructure"
	pb "github.com/echo-marche/hack-tech-tips-api/proto/pb/main_api"
	"github.com/echo-marche/hack-tech-tips-api/servers"
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
	sampleServer := &servers.SampleServer{}
	pb.RegisterMainApiServer(server, sampleServer)
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
