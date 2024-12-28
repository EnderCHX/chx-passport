package cmd

import (
	"chx-passport/api"
	"chx-passport/setup"
)

func main() {
	setup.Init()
	api.RunApi()
}
