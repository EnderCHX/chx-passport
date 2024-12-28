package setup

import (
	"chx-passport/config"
	"chx-passport/database"
	"chx-passport/models/user"
)

func Init() {
	config.Init()     //初始化配置
	database.InitDB() //初始化数据库
	user.InitTable()  //初始化用户表
	user.InitAdmin()  //初始化管理员
}
