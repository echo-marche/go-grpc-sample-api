package main

import (
	"github.com/echo-marche/niche-farm-bff/infrastructure"
)

func main() {
	// Presence API connection
	presenceApiConn := infrastructure.NewPresenceApiConnector()
	defer presenceApiConn.Close()
	// Article API connection
	articleApiConn := infrastructure.NewArticleApiConnector()
	defer articleApiConn.Close()
	// API settings
	router := infrastructure.Router{
		PresenceApiConn: presenceApiConn,
		ArticleApiConn:  articleApiConn,
	}
	router.InitRouter()
	// Start API Server
	router.Start()
}
