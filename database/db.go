package database

import (
	"chx_passport/config"
	"chx_passport/user"
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

func Init() {
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
			log.Println("数据库" + config.ConfigContext.MySQLConfig.DBName + "不存在，创建数据库")
			createMySQL, err := gorm.Open(mysql.Open(config.ConfigContext.MySQLConfig.Username+":"+
				config.ConfigContext.MySQLConfig.Password+"@tcp("+config.ConfigContext.MySQLConfig.Host+":"+
				config.ConfigContext.MySQLConfig.Port+")/?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
			//连接数据库，不带数据库名
			if err != nil {
				log.Println(err)
			}

			err = createMySQL.Exec("CREATE DATABASE IF NOT EXISTS " + config.ConfigContext.MySQLConfig.DBName).Error
			//创建数据库
			if err != nil {
				log.Println(err)
			}

			err = createMySQL.AutoMigrate(&user.User{})
			//创建表
			if err != nil {
				log.Println(err)
			}
			closeDB, _ := createMySQL.DB()
			closeDB.Close()
			//关闭数据库连接

			log.Println("数据库" + config.ConfigContext.MySQLConfig.DBName + "创建成功，重新连接数据库") //重试连接数据库
			MySQL, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
			if err != nil {
				log.Println(err)
			}
			log.Println("数据库连接成功")
		} else {
			log.Println(err)
			return
		}
	}
	log.Println("数据库连接成功")

	err = MySQL.AutoMigrate(&user.User{})
	//创建表
	if err != nil {
		log.Println(err)
	}

	closeMySQL, err := MySQL.DB()
	if err != nil {
		log.Println(err)
		return
	}
	defer closeMySQL.Close() //关闭数据库连接

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.ConfigContext.RedisConfig.Host + ":" + config.ConfigContext.RedisConfig.Port,
		Password: config.ConfigContext.RedisConfig.Password,
		DB:       config.ConfigContext.RedisConfig.DB,
	})
	log.Println("Redis连接成功")
	defer Rdb.Close() //关闭redis连接
}

func CreateUser(u *user.User) error {
	return MySQL.Create(u).Error
}
