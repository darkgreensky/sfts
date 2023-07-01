package controller

import (
	"net/http"
	"sfts/model"
	"sfts/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context) {
	leader := c.PostForm("leader")
	title := c.PostForm("title")
	intro := c.PostForm("introduction")
	time := c.PostForm("end_time")

	team := model.Team{
		Leader:       leader,
		Title:        title,
		Introduction: intro,
		EndTime:      time,
		Count:        1,
	}

	err := service.CreateTeam(team)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "创建队伍失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		Result:     "success",
		StatusMsg:  "创建队伍成功！",
	})
}

func SearchTeam(c *gin.Context) {
	team, err := service.SearchTeam()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取队伍列表失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetTeams{
		Response: model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "获取队伍列表成功",
		},
		Teams: team,
	})
}

func GetTeamINfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	team, err := service.GetTeamById(id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息失败:" + err.Error(),
		})
		return
	}

	member, err := service.GetTeamMember(id)

	c.JSON(http.StatusOK, model.GetTeam{
		Response: model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "获取信息成功",
		},
		Teams:   team,
		Members: member,
	})
}

func GetOwnTeam(c *gin.Context) {
	leader := c.Query("leader")

	team, err := service.GetTeamsByUser(leader)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取队伍失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetTeams{
		Response: model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "获取队伍成功",
		},
		Teams: team,
	})
}

func AddToTeam(c *gin.Context) {
	user := c.PostForm("user")
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	apply := model.Apply{
		UserName: user,
		TeamId:   id,
	}
	err = service.AddToTeam(apply)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "加入队伍失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		Result:     "success",
		StatusMsg:  "加入队伍成功！",
	})
}

func RemoveFromTeam(c *gin.Context) {
	user := c.PostForm("user")
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}
	err = service.RemoveFromTeam(user, id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "退出队伍失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		Result:     "success",
		StatusMsg:  "退出队伍成功！",
	})
}

func TeamCheck(c *gin.Context) {
	user := c.Query("user")
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取ID错误:" + err.Error(),
		})
		return
	}

	check, err := service.CheckAddTeam(user, id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "查询组队错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.CheckTeam{
		Response: model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "查询组队成功！",
		},
		Check: check,
	})
}
