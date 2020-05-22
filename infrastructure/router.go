package infrastructure

import (
	"fmt"
	"os"

	"github.com/echo-marche/hack-tech-tips-api/config"
	"github.com/echo-marche/hack-tech-tips-api/controllers"
	articlePb "github.com/echo-marche/hack-tech-tips-api/proto/pb/article"
	presencePb "github.com/echo-marche/hack-tech-tips-api/proto/pb/presence"
	sendmailPb "github.com/echo-marche/hack-tech-tips-api/proto/pb/sendmail"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
)

type Router struct {
	e               *echo.Echo
	PresenceApiConn *grpc.ClientConn
	ArticleApiConn  *grpc.ClientConn
	SendmailApiConn *grpc.ClientConn
}

func (router *Router) InitRouter() {
	// Echo instance
	router.e = echo.New()

	// Generate PB Client
	presenceClient := presencePb.NewPresenceClient(router.PresenceApiConn)
	articleClient := articlePb.NewArticleClient(router.ArticleApiConn)
	sendmailClient := sendmailPb.NewSendmailClient(router.SendmailApiConn)

	// Controller settings
	sampleController := controllers.SampleController{
		SendmailClient: sendmailClient,
	}
	userController := controllers.UserController{
		PresenceClient: presenceClient,
	}
	articleController := controllers.ArticleController{
		ArticleClient: articleClient,
	}
	healthzController := controllers.HealthzController{
		PresenceClient: presenceClient,
	}

	// Middleware
	router.e.Use(middleware.Logger())
	router.e.Use(middleware.Recover())
	if config.IsDev() {
		router.
			e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody))
		}))
	}

	// sample
	router.e.GET("/sample/sendmail", func(c echo.Context) error { return sampleController.SendMail(c) })
	// helthz
	router.e.GET("/healthz", func(c echo.Context) error { return healthzController.Index(c) })
	// ユーザー関連
	router.e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	router.e.POST("/user/temp_registration", func(c echo.Context) error { return userController.TempRegistration(c) })
	// 記事関連
	router.e.GET("/articles", func(c echo.Context) error { return articleController.Index(c) })
	// e.GET("/article/:id", func(c echo.Context) error { return articleController.Show(c) })
	// e.POST("/create", func(c echo.Context) error { return articleController.Create(c) })
	// e.PUT("/article/:id", func(c echo.Context) error { return articleController.Save(c) })
	// e.DELETE("/article/:id", func(c echo.Context) error { return articleController.Delete(c) })
}

func (router *Router) Start() {
	router.e.Logger.Fatal(router.e.Start(":" + config.GetEnv("MAIN_API_SERVICE_PORT")))
}
