package user

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	CustomConfig string `json:"custom_config"`
}

type AdminList struct {
	ID int `json:"id"`
}

// func Init() {
// 	admin := User{
// 		ID:           1,
// 		Username:     "admin",
// 		Password:     "admin",
// 		Email:        "admin@admin.com",
// 		CustomConfig: "",
// 	}
// 	log.Println(admin)
// 	database.MySQL.Create(&admin)
// }
