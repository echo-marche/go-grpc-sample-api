package controllers

import (
	"context"
	"fmt"
	"net/http"

	presencePb "github.com/echo-marche/hack-tech-tips-api/proto/pb/presence"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthzController struct {
	PresenceClient presencePb.PresenceClient
	// ArticleClient  articlePb.ArticleClient
}

func (controller *HealthzController) Index(c Context) (err error) {
	presenceRes, err := controller.PresenceClient.Healthz(context.TODO(), &emptypb.Empty{})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	fmt.Println(presenceRes)

	return c.JSON(http.StatusOK, "Haelthz OK!!")
}
