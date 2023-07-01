package router

import (
	"net/http"
	"sfts/controller"

	"github.com/gin-gonic/gin"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			//主要设置Access-Control-Allow-Origin
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func InitRouter() {
	r := gin.Default()
	r.Use(logger())

	r.GET("/ticket/per", controller.GetPerInformation)
	r.GET("/ticket/all", controller.GetAllInformation)
	r.GET("/ticket/search", controller.SearchInformation)
	r.POST("/ticket/buy", controller.BuyTicket)

	r.GET("/user/register", controller.Register)
	r.GET("/user/login", controller.Login)

	r.GET("/money/query", controller.QueryCurrency)
	r.POST("/money/recharge", controller.AddCurrency)

	r.POST("/image", controller.CreateImage)
	r.POST("/guide/create", controller.CreateGuide)
	r.POST("/guide/delete", controller.RemoveGuide)
	r.GET("/guide/all", controller.GetTitle)
	r.GET("/guide/per", controller.GetContent)
	r.GET("/guide/personal", controller.GetPersonal)

	r.POST("/team/create", controller.CreateTeam)
	r.GET("/team/all", controller.SearchTeam)
	r.GET("/team/per", controller.GetTeamINfo)
	r.GET("/team/personal", controller.GetOwnTeam)
	r.GET("/team/check", controller.TeamCheck)

	r.POST("/team/add", controller.AddToTeam)
	r.POST("/team/exit", controller.RemoveFromTeam)

	r.POST("/comment/create", controller.CreateComment)
	r.GET("/comment/per", controller.GetComment)

	r.GET("/program/all/time", controller.GetAllProgramsByTime)
	r.GET("/program/all/read", controller.GetAllProgramsByCount)
	r.GET("/program/per", controller.GetProgramInfo)

    r.Static("/static", "./static") // 设置静态文件目录

    r.NoRoute(func(c *gin.Context) {
		c.File("./index.html")
	})

	r.Run(":8080")
}
