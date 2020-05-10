package main

import (
	"github.com/echo-marche/hack-tech-tips-api/infrastructure"
)

func main() {
	// Presence API connection
	presenceApiConn := infrastructure.NewPresenceApiConnector()
	defer presenceApiConn.Close()
	// Article API connection
	articleApiConn := infrastructure.NewArticleApiConnector()
	defer articleApiConn.Close()
	// sendmail API connection
	sendmailApiConn := infrastructure.NewSendmailApiConnector()
	defer sendmailApiConn.Close()
	// API settings
	router := infrastructure.Router{
		PresenceApiConn: presenceApiConn,
		ArticleApiConn:  articleApiConn,
		SendmailApiConn: sendmailApiConn,
	}
	router.InitRouter()
	// Start API Server
	router.Start()
}
