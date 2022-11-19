package http

import (
	"exercise/domain/entity"
	"exercise/domain/web"
	"exercise/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserUsecase entity.UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase entity.UserUsecase) {
	handler := &UserHandler{UserUsecase: userUsecase}
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}

func (u *UserHandler) Register(c *gin.Context) {
	var userRegister entity.UserRegister
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	token, err := u.UserUsecase.Register(c.Request.Context(), &userRegister)
	if err != nil {
		c.JSON(helper.GetStatusCode(err), web.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, web.ResponseSuccess{Token: token})
}

func (u *UserHandler) Login(c *gin.Context) {
	var userLogin entity.UserLogin
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, web.ResponseError{Message: "BAD REQUEST"})
		return
	}

	token, err := u.UserUsecase.Login(c.Request.Context(), &userLogin)
	if err != nil {
		c.JSON(helper.GetStatusCode(err), web.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, web.ResponseSuccess{Token: token})
}
