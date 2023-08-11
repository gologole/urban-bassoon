package users

import (
	"fmt"
	"forummodule/sqllite"
	"gorm.io/gorm"
)

type ServiceLoginInput struct {
	Input sqllite.LoginInput
}

type ServiceUserModel struct {
	Usermodel sqllite.UserModel
}

func CreateDB() gorm.DB {
	db := sqllite.CreateUsersDB()
	return db
}

func ServiceLogin(db gorm.DB, UserInput sqllite.LoginInput) (sqllite.UserModel, error) {
	fmt.Println(UserInput)
	//fmt.Println("login ", login)
	//fmt.Println(" pass ", password)
	model, err := sqllite.Login(db, UserInput.Login, UserInput.Password)
	if err != nil {
		return sqllite.UserModel{}, err
	}
	return model, nil
}

func ServiceRegistrationUser(db gorm.DB, data sqllite.LoginInput) {

	sqllite.CreateUser(db, data.Login, data.Password)
}
