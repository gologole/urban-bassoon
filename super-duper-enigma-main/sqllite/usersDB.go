package sqllite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserModel struct {
	ID             int    `gorm:"primaryKey"`     // (уникальный идентификатор пользователя)
	Email          string `gorm:"email"`          // (электронная почта пользователя)
	Username       string `gorm:"username"`       // (логин пользователя)
	Password       string `gorm:"password"`       // (хэш пароля пользователя)
	AvatarURL      string `gorm:"avatar"`         // (ссылка на аватар пользователя)
	CreatedAt      string `gorm:"datacreating"`   // (дата создания аккаунта)
	RatingLikes    string `gorm:"RatingLikes"`    // (количество лайков, полученных пользователями за его сообщения и треды)
	RatingDislikes string `gorm:"RatingDisLikes"` // (количество дизлайков, полученных пользователями за его сообщения и треды)
}

func CreateUsersDB() gorm.DB {
	db, err := gorm.Open(sqlite.Open("users.db"))
	if err != nil {
		log.Println("table cant be open")
	}
	er := db.AutoMigrate(&UserModel{})
	if er != nil {
		log.Println("startserver.go 19 line,err=", er)
	}
	return *db
}

func NewUserModel(id int, email string, username string, password string, avatarURL string, createdAt string, ratingLikes string, ratingDislikes string) *UserModel {
	return &UserModel{
		ID:             id,
		Email:          email,
		Username:       username,
		Password:       password,
		AvatarURL:      avatarURL,
		CreatedAt:      createdAt,
		RatingLikes:    ratingLikes,
		RatingDislikes: ratingDislikes,
	}
}

func CreateUser(db gorm.DB, login, password string) {
	t := time.Now()
	cur := t.String()
	var count int64
	db.Model(&UserModel{}).Count(&count)
	newID := int(count + 1)

	setuser := NewUserModel(
		newID,
		"email",
		login,
		password,
		"avatar",
		cur,
		"0",
		"0",
	)
	SetUser(db, UserModel(*setuser))
}

// создание пользователя
func SetUser(db gorm.DB, user UserModel) {

	db.Create(&user)

	db.Save(&user)

	fmt.Println("User created successfully")
	fmt.Println(user)
}

// Функция для поиска пользователя по логину и паролю
func Login(db gorm.DB, username string, password string) (UserModel, error) {
	user := UserModel{}

	// Находим первого пользователя, у которого логин и пароль совпадают с переданными значениями
	// в душе не чаю заглавные ли здесь должны быть буквы
	err := db.Where("Username = ? AND Password = ?", username, password).First(&user).Error
	if err != nil {
		return UserModel{}, err // возращаем пустую структуру и ошибку, если пользователь не найден
	}
	return user, nil // возвращаем структуру пользователя
}
