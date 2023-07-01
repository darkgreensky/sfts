package controller

import (
	"net/http"
	"sfts/model"
	"sfts/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	user := c.PostForm("user")
	contend := c.PostForm("content")
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	comment := model.Comment{
		UserName: user,
		Content:  contend,
		TeamId:   id,
	}

	err = service.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "创建评论失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		Result:     "success",
		StatusMsg:  "创建评论成功！",
	})
}

func GetComment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	comment, err := service.GetComment(id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取评论失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.GetComments{
		Response: model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "获取评论成功！",
		},
		Comments: comment,
	})
}
