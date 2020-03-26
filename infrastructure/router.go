package infrastructure

import (
	"github.com/echo-marche/nicheye-bff/config"
	"github.com/echo-marche/nicheye-bff/controllers"
	articlePb "github.com/echo-marche/nicheye-bff/pb/article"
	presencePb "github.com/echo-marche/nicheye-bff/pb/presence"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
)

type Router struct {
	e               *echo.Echo
	PresenceApiConn *grpc.ClientConn
	ArticleApiConn  *grpc.ClientConn
}

func (router *Router) InitRouter() {
	// Echo instance
	router.e = echo.New()

	// Generate PB Client
	presenceClient := presencePb.NewPresenceClient(router.PresenceApiConn)
	articleClient := articlePb.NewArticleClient(router.ArticleApiConn)

	// Controller settings
	userController := controllers.UserController{
		PresenceClient: presenceClient,
	}
	articleController := controllers.ArticleController{
		ArticleClient: articleClient,
	}

	// Middleware
	router.e.Use(middleware.Logger())
	router.e.Use(middleware.Recover())

	// ユーザー関連
	router.e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	// 記事関連
	router.e.GET("/articles", func(c echo.Context) error { return articleController.Index(c) })
	// e.GET("/article/:id", func(c echo.Context) error { return articleController.Show(c) })
	// e.POST("/create", func(c echo.Context) error { return articleController.Create(c) })
	// e.PUT("/article/:id", func(c echo.Context) error { return articleController.Save(c) })
	// e.DELETE("/article/:id", func(c echo.Context) error { return articleController.Delete(c) })
}

func (router *Router) Start() {
	router.e.Logger.Fatal(router.e.Start(":" + config.GetEnv("API_PORT")))
}
