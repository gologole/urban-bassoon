package app

import (
	"fmt"
	"forummodule/delivery/http"
	"forummodule/service/users"
	"forummodule/sqllite"
	"gorm.io/gorm"
)

func Run() {
	db := users.CreateDB()
	go FisrstUser(db)

	StartServer(db)

}

func StartServer(db gorm.DB) {
	http.Handlereg(db)

}

func FisrstUser(db gorm.DB) {
	test := sqllite.LoginInput{
		"admin",
		"admin",
	}
	err := users.ServiceRegistrationUser(db, test)

	if err != nil {
		fmt.Println("Не удалось зарегистрировать пользователя", err)
	}
}
