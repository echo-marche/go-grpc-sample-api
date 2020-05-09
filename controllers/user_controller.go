package controllers

import (
	"context"
	"fmt"
	"net/http"

	presencePb "github.com/echo-marche/hack-tech-tips-api/proto/pb/presence"
	"github.com/echo-marche/hack-tech-tips-api/validations"
)

type UserController struct {
	PresenceClient presencePb.PresenceClient
}

func (controller *UserController) Index(c Context) (err error) {
	request := &presencePb.UserListRequest{StatusCode: ""}
	res, err := controller.PresenceClient.GetUserList(context.TODO(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	return c.JSON(http.StatusOK, res)
}

// ユーザー仮登録
func (controller *UserController) TempRegistration(c Context) (err error) {
	fmt.Println(c)

	request := &presencePb.UserRegistrationRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	// validation
	err = validations.ValidateUserRegistrationRequest(request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	res, err := controller.PresenceClient.UserRegistration(context.TODO(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	return c.JSON(http.StatusOK, res)
}

// func (controller *UserController) Show(c Context) (err error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user, err := controller.Db.UserById(id)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// 	return
// }

// func (controller *UserController) Create(c Context) (err error) {
// 	u := domain.User{}
// 	c.Bind(&u)
// 	user, err := controller.Db.Add(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, user)
// 	return
// }

// func (controller *UserController) Save(c Context) (err error) {
// 	u := domain.User{}
// 	c.Bind(&u)
// 	user, err := controller.Db.Update(u)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(201, user)
// 	return
// }

// func (controller *UserController) Delete(c Context) (err error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user := domain.User{
// 		ID: id,
// 	}
// 	err = controller.Db.DeleteById(user)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// 	return
// }
