package user

import (
	"chx-passport/config"
	"chx-passport/database"
	"log"
	"time"

	tools "github.com/EnderCHX/chx-tools-go/encrypt"
)

type User struct {
	Username          string `json:"username" gorm:"primaryKey"`
	Password          string `json:"password"`
	passwordEncrypted bool
	Email             string    `json:"email" gorm:"unique;index"`
	Role              string    `json:"role"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
	Deleted           bool      `json:"deleted"`
	CustomConfig      string    `json:"custom_config"`
}

type UserReqBody struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	ChangePwdOld string `json:"change_pwd_old"`
	ChangePwdNew string `json:"change_pwd_new"`
}

var (
	RoleList = []string{"admin", "user"}
)

func NewUser(username, password, email string, role int) *User {
	return &User{
		Username:          username,
		Password:          tools.Sha256(password),
		passwordEncrypted: false,
		Email:             email,
		Role:              RoleList[role],
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Deleted:           false,
	}
}

func NewUserReqBody(username, password, email string) *UserReqBody {
	return &UserReqBody{
		Username: username,
		Password: password,
		Email:    email,
	}
}

func (u *UserReqBody) ToUser() *User {
	return &User{
		Username:          u.Username,
		Password:          u.Password,
		Email:             u.Email,
		passwordEncrypted: false,
	}
}
func InitTable() {
	err := database.MySQL.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}
}

func InitAdmin() {
	admin := User{
		Username:          "admin",
		Password:          tools.Sha256("admin"),
		passwordEncrypted: false,
		Email:             "admin@chxc.cc",
		Role:              "admin",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Deleted:           false,
		CustomConfig:      "",
	}
	log.Println(admin)
	err := admin.Insert()
	if err != nil {
		log.Println(err)
	}
}

func (u *User) PasswordEncrypt() *User {
	u.Password = tools.Sha256(u.Password + config.ConfigContext.SecretKeys.PasswdSecret)
	u.passwordEncrypted = true
	return u
}
func (u *User) Insert() error {
	if !u.passwordEncrypted {
		u.PasswordEncrypt()
	}
	return database.MySQL.Create(u).Error
}

func (u *User) PasswordCheck() bool {
	_u := &User{}
	err := database.MySQL.Select("password", "deleted").Where("username = ?", u.Username).First(_u).Error
	if err != nil {
		return false
	}
	if !u.passwordEncrypted {
		u.PasswordEncrypt()
	}
	return _u.Password == u.Password
}
func (u *User) Update() error {
	if !u.passwordEncrypted {
		u.PasswordEncrypt()
	}
	return database.MySQL.Save(u).Error
}
func (u *User) Delete() error {
	u.Deleted = true
	u.DeletedAt = time.Now()
	return database.MySQL.Save(u).Error
}

func (u *User) SelectEmail() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("email").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}

func (u *User) SelectRole() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("role").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}

func (u *User) SelectPassword() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("password").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}

func (u *User) SelectCreatedAt() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("created_at").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}
func (u *User) SelectUpdatedAt() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("updated_at").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}

func (u *User) SelectDeletedAt() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("deleted_at").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}

func (u *User) SelectDeleted() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("deleted").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}

func (u *User) SelectCustomConfig() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Select("custom_config").Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}

func (u *User) SelectAll() *User {
	if u.Username == "" {
		return u
	}
	err := database.MySQL.Where("username = ?", u.Username).First(u).Error
	if err != nil {
		log.Println(err)
	}
	u.passwordEncrypted = true
	return u
}
