package middleware

import (
	"chx-passport/auth"
	"chx-passport/config"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var url string = config.ConfigContext.WebUrl

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, url+"/login")
			c.Abort()
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := auth.VerifyToken(token, config.ConfigContext.Secret)
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
