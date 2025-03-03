package middleware

import (
	"chx-passport/auth"
	"chx-passport/config"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":  "Authorization header is empty",
				"code": "InvalidAuthorizationHeader",
				"data": nil,
			})
			c.Abort()
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := auth.VerifyToken(token, config.ConfigContext.SecretKeys.AccessSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":  err,
				"code": "InvalidAccessToken",
				"data": nil,
			})
			c.Abort()
			return
		} else {
			c.Set("claims", claims)
			c.Next()
		}
	}
}

func ShowUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			log.Println("User : Guest", c.Request.RemoteAddr, "from", c.Request.RequestURI, "is visiting", c.Request.URL.Path)
		} else {
			token = strings.Replace(token, "Bearer ", "", 1)
			claims, err := auth.VerifyToken(token, config.ConfigContext.SecretKeys.AccessSecret)
			if err != nil {
				log.Println("User : Guest", c.Request.RemoteAddr, "from", c.Request.RequestURI, "is visiting", c.Request.URL.Path)
			} else {
				log.Println("User : ", claims.Username, c.Request.RemoteAddr, "from", c.Request.RequestURI, "is visiting", c.Request.URL.Path)
			}
		}

		c.Next()
	}
}

func Cors() gin.HandlerFunc { //跨域中间件
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
