package controller

import (
	"chx-passport/database"
	"chx-passport/models/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	_user := user.User{}
	c.BindJSON(&_user)
	_user.CreatedAt = time.Now()
	_user.UpdatedAt = time.Now()
	_user.Role = user.RoleList[0]
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
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data":    _user,
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
