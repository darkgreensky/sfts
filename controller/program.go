package controller

import (
	"net/http"
	"sfts/model"
	"sfts/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProgramsByTime(c *gin.Context) {
	program, err := service.GetProgramByTime()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetPrograms{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功！",
		},
		Programs: program,
	})
}

func GetAllProgramsByCount(c *gin.Context) {
	program, err := service.GetProgramByCount()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetPrograms{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功！",
		},
		Programs: program,
	})
}

func GetProgramInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	program, err := service.GetProgramById(id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetProgram{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功！",
		},
		Programs: program,
	})
}
