package models

import (
	"github.com/dgquijote/be-screening/auth"
	"github.com/dgquijote/be-screening/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	IsSeller bool   `json:"is_seller"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func GetUserByToken(tokenString string) (User, *gorm.DB) {
	var user User

	tokenUser := auth.TokenUser(tokenString)

	record := database.Instance.Where("username = ?", tokenUser).First(&user)

	return user, record
}
