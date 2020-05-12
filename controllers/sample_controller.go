package controllers

import (
	"context"
	"fmt"
	"net/http"

	sendmailPb "github.com/echo-marche/hack-tech-tips-api/proto/pb/sendmail"
)

type SampleController struct {
	SendmailClient sendmailPb.SendmailClient
}

func (controller *SampleController) Index(c Context) (err error) {
	request := &sendmailPb.EmailRequest{
		FromUserName: "テストユーザー",
		FromAddress:  "htt_test@htt.support.com",
		ToAddress:    "to_test@gmail.com",
		Subject:      "テスト件名",
		Msg:          "テスト本文テスト本文テスト本文テスト本文テスト本文",
	}

	res, err := controller.SendmailClient.SendSample(context.TODO(), request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	return c.JSON(http.StatusOK, res)
}
