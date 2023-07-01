package controller

import (
	"net/http"
	"sfts/model"
	"sfts/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateImage(c *gin.Context) {
	data, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取图片失败:" + err.Error(),
		})
		return
	}

	imageURL, err := service.UploadImage(data)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "上传图片失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.ImageResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取图片成功！",
		},
		ImageURL: imageURL,
	})
}

func CreateGuide(c *gin.Context) {
	author := c.PostForm("author")
	title := c.PostForm("title")
	content := c.PostForm("content")

    if content == "" {
        content = "无内容"
    }

    if title == "" {
        title = "无标题"
    }

	guide := model.Guide{
		Author:  author,
		Content: content,
		Title:   title,
	}

	err := service.CreateGuide(guide)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "创建路线推荐失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "创建路线推荐成功！",
	})
}

func GetTitle(c *gin.Context) {
	guide, err := service.GetTitle()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取攻略分享失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetGuides{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取攻略分享成功",
		},
		Guides: guide,
	})
}

func GetContent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	guide, err := service.GetGuideById(id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息失败:" + err.Error(),
		})
		return
	}

	service.ReadCount(id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "更新阅读量失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetGuide{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取信息成功",
		},
		Guides: guide,
	})
}

func RemoveGuide(c *gin.Context) {
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取id失败:" + err.Error(),
		})
		return
	}

	err = service.RemoveGuide(id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "删除攻略分享失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		Result:     "success",
		StatusMsg:  "删除攻略分享成功！",
	})
}

func GetPersonal(c *gin.Context) {
	author := c.Query("author")

	guide, err := service.GetGuideByAuthor(author)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取攻略失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetGuides{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取攻略分享成功",
		},
		Guides: guide,
	})
}
