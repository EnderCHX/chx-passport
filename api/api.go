package api

import (
	"chx_passport/config"
	"log"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	gin.SetMode(config.ConfigContext.ApiConfig.Mode)
	host := config.ConfigContext.ApiConfig.Host
	port := config.ConfigContext.ApiConfig.Port
	log.Println("Starting API on http://" + host + ":" + port)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(host + ":" + port)
}
