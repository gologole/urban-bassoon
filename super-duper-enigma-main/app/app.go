package app

import (
	"fmt"
	"forummodule/delivery/http"
	"forummodule/service/users"
	"forummodule/sqllite"
	"log"
)

func Run() {

	db := users.CreateDB()

	test := users.ServiceLoginInput{
		sqllite.LoginInput{
			"admin1",
			"admin1",
		},
	}

	test2 := users.ServiceLoginInput{
		sqllite.LoginInput{
			"admin2",
			"admin2",
		},
	}
	users.ServiceRegistrationUser(db, test)

	model, err := users.ServiceLogin(db, test2)
	if err != nil {
		log.Println("Пользователя не существует")
	}

	fmt.Print(model)

	//go http.Handlereg()

}

func StartServer() {
	go http.Handlereg()
}
