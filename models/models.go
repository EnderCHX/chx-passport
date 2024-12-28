package models

import (
	"chx-passport/models/user"
)

type modelsList struct {
	User user.User `json:"user"`
}

var ModelsList modelsList

func init() {
	ModelsList := modelsList{
		User: user.User{},
	}
}
