package controller

import (
	"net/http"
	"sfts/model"
	"sfts/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if len(password) < 4 {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "密码长度不足4，请重新设置",
		})
		return
	}

	node, err := service.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "Register Error" + err.Error(),
		})
		return
	}

	if node == 0 {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "注册成功",
		})
		return
	} else {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户名已被注册，请重新设置",
		})
		return
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	node, err := service.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "Login Error" + err.Error(),
		})
	}

	if node == 0 {
		c.JSON(http.StatusOK, model.UserResponse{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "登录成功",
			Token:      username,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户名或密码错误，请重新登陆",
		})
	}
}
