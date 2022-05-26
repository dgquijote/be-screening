package models

import (
	"log"

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

func MigrateUsers() {
	result := map[string]interface{}{}

	record := database.Instance.Model(&User{}).First(&result)

	if record.Error != nil {
		database.Instance.AutoMigrate(&User{})

		var users = []User{
			{
				Email:    "test.seller@email.com",
				Username: "test.seller",
				Password: "123456789",
				Name:     "Test Seller",
				IsSeller: true,
			},
			{
				Email:    "test.seller2@email.com",
				Username: "test.seller2",
				Password: "123456789",
				Name:     "Test Seller 2",
				IsSeller: true,
			},
			{
				Email:    "test.user@email.com",
				Username: "test.user",
				Password: "123456789",
				Name:     "Test User",
				IsSeller: false,
			},
			{
				Email:    "test.user2@email.com",
				Username: "test.user2",
				Password: "123456789",
				Name:     "Test User 2",
				IsSeller: false,
			},
		}

		for i, u := range users {
			users[i].HashPassword(u.Password)
		}

		database.Instance.Create(&users)

		log.Println("Database Users Migration Completed!")
	}
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
