package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to Database!")
	}
	log.Println("Connected to Database!")
}

func MockConnect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}))
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to Database!")
	}
	log.Println("Connected to Database!")
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	IsSeller bool   `json:"is_seller"`
}

func Migrate() {
	Instance.AutoMigrate(&User{})

	seller := User{
		Email:    "test.seller@email.com",
		Username: "test.seller",
		Password: "123456789",
		Name:     "Test Seller",
		IsSeller: true,
	}

	seller.HashPassword(seller.Password)

	Instance.Create(&seller)

	seller2 := User{
		Email:    "test.seller2@email.com",
		Username: "test.seller2",
		Password: "123456789",
		Name:     "Test Seller 2",
		IsSeller: true,
	}

	seller2.HashPassword(seller2.Password)

	Instance.Create(&seller2)

	user := User{
		Email:    "test.user@email.com",
		Username: "test.user",
		Password: "123456789",
		Name:     "Test User",
		IsSeller: false,
	}

	user.HashPassword(user.Password)

	Instance.Create(&user)

	user2 := User{
		Email:    "test.user2@email.com",
		Username: "test.user2",
		Password: "123456789",
		Name:     "Test User 2",
		IsSeller: false,
	}

	user2.HashPassword(user2.Password)

	Instance.Create(&user2)

	Instance.AutoMigrate(&User{})

	log.Println("Database Migration Completed!")
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
