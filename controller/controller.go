package controller

import (
	"chx-passport/auth"
	"chx-passport/config"
	"chx-passport/models/user"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userreqbody := user.UserReqBody{}
	c.BindJSON(&userreqbody)

	if userreqbody.Username == "" || userreqbody.Password == "" || userreqbody.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名、密码、邮箱不能为空",
			"data": nil,
			"code": "InvalidParameter",
		})
		return
	}

	pattern := "^[a-zA-Z0-9_]+$"
	match, _ := regexp.MatchString(pattern, userreqbody.Username)

	if !match {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名只能包含字母、数字和下划线",
			"code": "InvalidUsername",
			"data": nil,
		})
		return
	}

	pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ = regexp.MatchString(pattern, userreqbody.Email)

	if !match {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "邮箱格式不正确",
			"code": "InvalidEmail",
			"data": nil,
		})
		return
	}

	if len(userreqbody.Password) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "密码长度不能小于6位",
			"code": "InvalidPassword",
			"data": nil,
		})
		return
	}

	_user := user.NewUser(
		userreqbody.Username,
		"",
		userreqbody.Email,
		user.RoleUser,
	)
	_user.Password = userreqbody.Password

	if userreqbody.Avatar == "" {
		_user.Avatar = "https://www.gravatar.com/avatar/" + strings.ToLower(strings.TrimSpace(userreqbody.Email))
	}

	if userreqbody.Signature == "" {
		_user.Signature = "系统原装签名！"
	}

	err := _user.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名或邮箱已存在",
			"code": "UserAlreadyExists",
			"data": nil,
		})
		return
	}

	refresh_token, _ := auth.GetToken(*_user, config.ConfigContext.SecretKeys.RefreshSecret, time.Hour*24*30)
	access_token, _ := auth.GetToken(*_user, config.ConfigContext.SecretKeys.AccessSecret, time.Hour)

	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"code": "Success",
		"data": gin.H{
			"refresh_token": refresh_token,
			"access_token":  access_token,
		},
	})
}

func Login(c *gin.Context) {
	uqb := user.UserReqBody{}
	c.BindJSON(&uqb)
	_user := uqb.ToUser()

	if _user.Username == "" || _user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名、密码不能为空",
			"code": "InvalidParameter",
			"data": nil,
		})
		return
	}

	if !_user.PasswordCheck() {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名或密码不正确",
			"code": "IncorrectUserNamweOrPassword",
			"data": nil,
		})
		return
	}

	_user.SelectRole().SelectAvatar().SelectSignature()

	refresh_token, _ := auth.GetToken(*_user, config.ConfigContext.SecretKeys.RefreshSecret, time.Hour*24*30)
	access_token, _ := auth.GetToken(*_user, config.ConfigContext.SecretKeys.AccessSecret, time.Hour)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"code": "Success",
		"data": gin.H{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
	})
}

func UserInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	_user := user.User{
		Username: claims.(*auth.JWTPayload).Username,
	}
	_user.SelectCreatedAt().SelectCustomConfig().SelectDeleted().SelectEmail().SelectRole().SelectAvatar().SelectSignature()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "User information retrieved successfully",
		"code": "Success",
		"data": _user,
	})
}

func ChangeInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	uqb := user.UserReqBody{}
	c.BindJSON(&uqb)

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, uqb.Email)

	if !match {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "邮箱格式不正确",
			"code": "InvalidEmail",
			"data": nil,
		})
		return
	}

	if len(uqb.ChangePwdNew) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "密码长度不能小于6位",
			"code": "InvalidPassword",
			"data": nil,
		})
		return
	}

	_user := user.User{
		Username:     claims.(*auth.JWTPayload).Username,
		Email:        uqb.Email,
		Password:     uqb.ChangePwdNew,
		CustomConfig: uqb.CustomConfig,
	}
	err := _user.Update()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "邮箱已存在",
			"code": "UserAlreadyExists",
			"data": nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "修改成功",
			"code": "Success",
			"data": nil,
		})
	}

}

type RefreshTokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

func RefreshToken(c *gin.Context) {
	rt := RefreshTokenReqBody{}
	c.Bind(&rt)

	if rt.RefreshToken == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Refresh token cannot be empty",
			"code": "InvalidParameter",
			"data": nil,
		})
		return
	}

	claims, err := auth.VerifyToken(rt.RefreshToken, config.ConfigContext.SecretKeys.RefreshSecret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err,
			"code": "InvalidRefreshToken",
			"data": nil,
		})
		return
	}
	_user := user.User{
		Username: claims.Username,
	}
	_user.SelectRole().SelectAvatar().SelectSignature()
	accessToken, err := auth.GetToken(_user, config.ConfigContext.SecretKeys.AccessSecret, time.Hour)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err,
			"code": "InternalServerError",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Refresh token successful",
		"code": "Success",
		"data": gin.H{
			"access_token": accessToken,
		},
	})
}

func VerifyAccessToken(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	accessToken = strings.Replace(accessToken, "Bearer ", "", 1)
	claims, err := auth.VerifyToken(accessToken, config.ConfigContext.SecretKeys.AccessSecret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err,
			"code": "InvalidAccessToken",
			"data": nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Access token verified",
			"code": "Success",
			"data": claims,
		})
	}
}
