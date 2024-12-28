package user

import (
	"chx-passport/database"
	"log"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CustomConfig string    `json:"custom_config"`
}

var (
	RoleList = []string{"admin", "user", "guest"}
)

func InitTable() {
	err := database.MySQL.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}
}

func InitAdmin() {
	admin := User{
		ID:           1,
		Username:     "admin",
		Password:     "admin",
		Email:        "admin@chxc.cc",
		Role:         "admin",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CustomConfig: "",
	}
	log.Println(admin)
	err := admin.Insert()
	if err != nil {
		log.Println(err)
	}
}

func (u *User) Insert() error {
	return database.MySQL.Create(u).Error
}
func (u *User) Update() error {
	return database.MySQL.Save(u).Error
}
func (u *User) Delete() error {
	return database.MySQL.Delete(u).Error
}
