package middleware

import (
	"chx-passport/auth"
	"chx-passport/config"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var url string = "https://chxc.cc"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, url+"/login")
			c.Abort()
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := auth.VerifyToken(token, config.ConfigContext.SecretKeys.RefreshSecret)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, url+"/")
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
		claims, ok := c.Get("claims")
		if !ok {
			log.Println("Claims not found in context")
			c.Next()
		} else {
			username := claims.(*auth.JWTPayload).Username
			log.Println("User :", username, c.Request.RemoteAddr, "from", c.Request.RequestURI, "is visiting", c.Request.URL.Path)
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
