package infrastructure

import (
	"log"

	"github.com/echo-marche/hack-tech-tips-api/config"
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

func NewSendmailApiConnector() *grpc.ClientConn {
	sendmailApiConfig := config.InitSendmailApiConfig()
	conn, err := grpc.Dial(sendmailApiConfig.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return conn
}

func NewPopularityApiConnector() *grpc.ClientConn {
	popularityApiConfig := config.InitPopularityApiConfig()
	conn, err := grpc.Dial(popularityApiConfig.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return conn
}

func NewDepositApiConnector() *grpc.ClientConn {
	depositApiConfig := config.InitDepositApiConfig()
	conn, err := grpc.Dial(depositApiConfig.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return conn
}

func NewNotificationApiConnector() *grpc.ClientConn {
	notificationApiConfig := config.InitNotificationApiConfig()
	conn, err := grpc.Dial(notificationApiConfig.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return conn
}

func NewInquiryApiConnector() *grpc.ClientConn {
	inquiryApiConfig := config.InitInquiryApiConfig()
	conn, err := grpc.Dial(inquiryApiConfig.Url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	return conn
}
