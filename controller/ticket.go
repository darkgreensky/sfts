package controller

import (
	"net/http"
	"sfts/model"
	"sfts/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPerInformation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	infor, err := service.GetInfoById(id)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetInformation{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功！",
		},
		Infor: infor,
	})
}

func GetAllInformation(c *gin.Context) {
	infor, err := service.GetAllInfo()

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetInformations{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功！",
		},
		Infor: infor,
	})
}

func SearchInformation(c *gin.Context) {
	text := c.Query("text")

	infor, err := service.SearchInfo(text)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetInformations{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功！",
		},
		Infor: infor,
	})
}
