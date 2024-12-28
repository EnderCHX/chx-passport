package api

import (
	"chx-passport/api/middleware"
	"chx-passport/config"
	"chx-passport/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	gin.SetMode(config.ConfigContext.ApiConfig.Mode)
	host := config.ConfigContext.ApiConfig.Host
	port := config.ConfigContext.ApiConfig.Port
	log.Println("Starting API on http://" + host + ":" + port)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Caillo World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	needLoginGroup := r.Group("/user/:username", middleware.Auth(), middleware.ShowUserInfo())
	{
		needLoginGroup.GET("/info", controller.UserInfo)
		// needLoginGroup.POST("/update", controller.UpdateUserInfo)
	}

	r.Run(host + ":" + port)
}
