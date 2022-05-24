package database

import (
	"log"

	"github.com/dgquijote/be-screening/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to Database!")
	}
	log.Println("Connected to Database!")
}

func MockConnect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to Database!")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})

	seller := models.User{
		Email:    "test.seller@email.com",
		Username: "test.seller",
		Password: "123456789",
		Name:     "Test Seller",
	}

	seller.HashPassword(seller.Password)

	Instance.Create(&seller)

	seller2 := models.User{
		Email:    "test.seller@email.com",
		Username: "test.seller2",
		Password: "123456789",
		Name:     "Test Seller 2",
	}

	seller2.HashPassword(seller2.Password)

	Instance.Create(&seller2)

	user := models.User{
		Email:    "test.user@email.com",
		Username: "test.user",
		Password: "123456789",
		Name:     "Test User",
	}

	user.HashPassword(user.Password)

	Instance.Create(&user)

	Instance.AutoMigrate(&models.Order{})

	Instance.AutoMigrate(&models.OrderLogistics{})

	log.Println("Database Migration Completed!")
}
