package app

import (
	"forummodule/delivery/http"
	"forummodule/service/users"
	"forummodule/sqllite"
	"gorm.io/gorm"
)

func Run() {

	db := users.CreateDB()

	test := sqllite.LoginInput{
		"admin",
		"admin",
	}

	users.ServiceRegistrationUser(db, test)

	http.Handlereg(db)

}

func StartServer(db gorm.DB) {
	http.Handlereg(db)
}
