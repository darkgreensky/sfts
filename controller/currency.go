package controller

import (
	"net/http"
	"sfts/model"
	"sfts/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryCurrency(c *gin.Context) {
	user := c.Query("user")

	money, err := service.QueryCurrency(user)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "查询余额错误:" + err.Error(),
		})
		return
	}

	ticket, err := service.GetTicketInfo(user)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取票务信息错误:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.GetMoney{
		Response: model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "查询票务信息成功！",
		},
		Money:  money,
		Ticket: ticket,
	})
}

func AddCurrency(c *gin.Context) {
	user := c.PostForm("user")
	money, err := strconv.ParseInt(c.PostForm("money"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取金额失败:" + err.Error(),
		})
		return
	}

	err = service.Recharge(user, money)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "充值失败:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		Result:     "success",
		StatusMsg:  "充值推荐成功！",
	})
}

func BuyTicket(c *gin.Context) {
	user := c.PostForm("user")
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			Result:     "false",
			StatusMsg:  "获取id失败:" + err.Error(),
		})
		return
	}

	status, rest, cost, err := service.BuyTicket(user, id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			Result:     "false",
			StatusMsg:  "购票失败:" + err.Error(),
		})
		return
	}

	if !status {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			Result:     "false",
			StatusMsg:  "购票失败，余额不足",
		})
		return
	}

	c.JSON(http.StatusOK, model.Buy{
		Response: model.Response{
			StatusCode: 0,
			Result:     "success",
			StatusMsg:  "购票成功！",
		},
		Money: rest,
		Cost:  cost,
	})
}
