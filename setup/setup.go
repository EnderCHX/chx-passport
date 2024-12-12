package setup

import (
	"chx_passport/config"
	"chx_passport/database"
)

func Init() {
	config.Init()
	database.Init()
}
