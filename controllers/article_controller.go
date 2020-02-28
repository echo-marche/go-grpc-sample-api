package controllers

import (
	"context"
	"net/http"

	articlePb "github.com/echo-marche/niche-farm-bff/pb/article"
)

type ArticleController struct {
	ArticleClient articlePb.ArticleClient
}

func (controller *ArticleController) Index(c Context) (err error) {
	request := &articlePb.ArticleListRequest{SystemCode: ""}
	res, err := controller.ArticleClient.GetArticleList(context.TODO(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	return c.JSON(http.StatusOK, res)
}

// func (controller *ArticleController) Show(c Context) (err error) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user, err := controller.Db.UserById(id)
// 	if err != nil {
// 		c.JSON(500, NewError(err))
// 		return
// 	}
// 	c.JSON(200, user)
// 	return
// }

// func (controller *ArticleController) Create(c Context) (err error) {
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

// func (controller *ArticleController) Save(c Context) (err error) {
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

// func (controller *ArticleController) Delete(c Context) (err error) {
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
