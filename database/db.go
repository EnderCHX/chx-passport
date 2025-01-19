package database

import (
	"chx-passport/config"
	"context"
	"errors"
	"log"

	mysqlError "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MySQL *gorm.DB
	Rdb   *redis.Client
	Dsn   string
)

func InitDB() {
	Dsn = config.ConfigContext.MySQLConfig.Username + ":" +
		config.ConfigContext.MySQLConfig.Password +
		"@tcp(" + config.ConfigContext.MySQLConfig.Host + ":" +
		config.ConfigContext.MySQLConfig.Port + ")/" +
		config.ConfigContext.MySQLConfig.DBName +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	MySQL, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		var mysqlErr *mysqlError.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1049 { //数据库不存在，创建数据库
			panic("数据库" + config.ConfigContext.MySQLConfig.DBName + "不存在，请手动创建数据库")
		}
		panic(err)
	}
	log.Println("数据库连接成功")

	// err = MySQL.AutoMigrate(&models.ModelsList.User)
	// //创建表
	// if err != nil {
	// 	log.Println(err)
	// }

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.ConfigContext.RedisConfig.Host + ":" + config.ConfigContext.RedisConfig.Port,
		Password: config.ConfigContext.RedisConfig.Password,
		DB:       config.ConfigContext.RedisConfig.DB,
	})
	ctx := context.Background()
	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Redis:", pong)
}
