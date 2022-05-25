package models

import (
	"net/http"

	"github.com/dgquijote/be-screening/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	SellerId   int
	Seller     User
	ReceiverId int
	Receiver   User
	Item       string `json:"item"`
	Weight     string `json:"weight"`
	ItemAmount string `json:"item_amount"`
}

func GetAll(c *gin.Context) []Order {
	tokenString := c.GetHeader("Authorization")

	user, record := GetUserByToken(tokenString)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return nil
	}
	var orders []Order
	if user.IsSeller {
		database.Instance.Preload("Seller").Preload("Receiver").Where("seller_id = ?", user.ID).Find(&orders)
	} else {
		database.Instance.Preload("Seller").Preload("Receiver").Where("receiver_id = ?", user.ID).Find(&orders)
	}

	return orders
}
