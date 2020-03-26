package infrastructure

import (
	"log"

	"github.com/echo-marche/nicheye-bff/config"
	"google.golang.org/grpc"
)

func NewPresenceApiConnector() *grpc.ClientConn {
	presenceApiConfig := config.InitPresenceApiConfig()
	conn, err := grpc.Dial(presenceApiConfig.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return conn
}

func NewArticleApiConnector() *grpc.ClientConn {
	articleApiConfig := config.InitArticleApiConfig()
	conn, err := grpc.Dial(articleApiConfig.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return conn
}
