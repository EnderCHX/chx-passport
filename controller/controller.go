package controller

import (
	"chx-passport/config"
	"chx-passport/database"
	"chx-passport/models/user"
	"go/token"
	"net/http"
	"time"

	tools "github.com/EnderCHX/chx-tools-go/encrypt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	_user := user.User{}
	c.BindJSON(&_user)
	_user.CreatedAt = time.Now()
	_user.UpdatedAt = time.Now()
	_user.Role = user.RoleList[0]
	_user.Password = tools.Sha256(_user.Password + config.ConfigContext.Secret)
	_user.Insert()
	_user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"data":    _user,
	})
}

func Login(c *gin.Context) {
	_user := user.User{}
	c.BindJSON(&_user)
	username := _user.Username
	password := tools.Sha256(_user.Password + config.ConfigContext.Secret)

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username or password cannot be empty",
		})
	}

	rs := database.MySQL.First(&_user)


	token := auth.GetToken(_user, config.ConfigContext.Secret)
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data":    _user,
		"token" : 
	})
}

func UserInfo(c *gin.Context) {
	username := c.Param("username")
	_user := user.User{}
	database.MySQL.Where("username = ?", username).First(&_user)
	_user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "User information retrieved successfully",
		"data":    _user,
	})
}
