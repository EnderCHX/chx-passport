package api

import (
	"chx-passport/api/middleware"
	"chx-passport/config"
	"chx-passport/controller"
	"chx-passport/utils/log"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	log.Setup("./log.log", "debug")
	logger := log.GetLogger()
	gin.SetMode(config.ConfigContext.ApiConfig.Mode)
	host := config.ConfigContext.ApiConfig.Host
	port := config.ConfigContext.ApiConfig.Port
	logger.Info("Starting API on http://" + host + ":" + port)
	r := gin.New()

	r.Use(log.GinZapLogger(), gin.Recovery(), middleware.Cors()) // 允许跨域

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

	r.POST("/refresh", controller.RefreshToken)

	needLoginGroup := r.Group("/user", middleware.Auth())
	{
		needLoginGroup.GET("/info", controller.UserInfo)
		needLoginGroup.POST("/change_info", controller.ChangeInfo)
		// needLoginGroup.POST("/update", controller.UpdateUserInfo)
	}

	r.Run(host + ":" + port)
}
