package main

import (
	"chx_passport/api"
	"chx_passport/setup"
)

func main() {
	setup.Init()
	api.RunApi()
}
