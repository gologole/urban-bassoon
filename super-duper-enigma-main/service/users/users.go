package users

import (
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

func ServiceLogin(db gorm.DB, UserInput ServiceLoginInput) (sqllite.UserModel, error) {
	login := UserInput.Input.Login
	password := UserInput.Input.Password
	model, err := sqllite.Login(db, login, password)
	if err != nil {
		return sqllite.UserModel{}, err
	}
	return model, nil
}

func ServiceRegistrationUser(db gorm.DB, data ServiceLoginInput) {
	login := data.Input.Login
	password := data.Input.Password
	sqllite.CreateUser(db, login, password)
}
