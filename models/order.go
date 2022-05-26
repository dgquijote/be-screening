package models

import (
	"log"
	"net/http"
	"time"

	"github.com/dgquijote/be-screening/database"
	"github.com/gin-gonic/gin"
)

type Order struct {
	Id         uint `gorm:"primary_key" json:"id"`
	SellerId   int  `json:"seller_id"`
	Seller     User
	ReceiverId int `json:"receiver_id"`
	Receiver   User
	Item       string    `json:"item"`
	Weight     string    `json:"weight"`
	ItemAmount string    `json:"item_amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func MigrateOrders() {
	result := map[string]interface{}{}

	record := database.Instance.Model(&Order{}).First(&result)

	if record.Error != nil {
		database.Instance.AutoMigrate(&Order{})

		date_ordered, _ := time.Parse(time.RFC3339, "2022-05-01T08:34:45Z")
		date_orderedu, _ := time.Parse(time.RFC3339, "2022-05-04T08:15:13Z")

		date_ordered2, _ := time.Parse(time.RFC3339, "2022-05-20T13:10:20Z")

		var records = []Order{
			{
				SellerId:   1,
				ReceiverId: 3,
				Item:       "Synology DiskStation DS718+ NAS Server for Business with Intel Celeron CPU, 6GB Memory, 8TB HDD Storage",
				ItemAmount: "999.00",
				Weight:     "3.84 lbs",
				CreatedAt:  date_ordered,
				UpdatedAt:  date_orderedu,
			},
			{
				SellerId:   2,
				ReceiverId: 3,
				Item:       "NVIDIA - GeForce RTX 3090 Ti - Titanium and black",
				ItemAmount: "1999.99",
				Weight:     "4.84 lbs",
				CreatedAt:  date_ordered2,
				UpdatedAt:  date_ordered2,
			},
		}

		database.Instance.Create(&records)

		log.Println("Database Orders Migration Completed!")
	}
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
		database.Instance.Where("seller_id = ?", user.ID).Find(&orders)
	} else {
		database.Instance.Where("receiver_id = ?", user.ID).Find(&orders)
	}

	return orders
}

func GetById(c *gin.Context) Order {
	tokenString := c.GetHeader("Authorization")

	user, record := GetUserByToken(tokenString)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return Order{}
	}
	var order Order

	id := c.Param("id")

	if user.IsSeller {
		database.Instance.Preload("Seller").Preload("Receiver").Where("id = ?", id).Where("seller_id = ?", user.ID).Find(&order)
	} else {
		database.Instance.Preload("Seller").Preload("Receiver").Where("id = ?", id).Where("receiver_id = ?", user.ID).Find(&order)
	}

	return order
}

func Add(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	user, record := GetUserByToken(tokenString)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	requestBody := Order{}

	c.BindJSON(&requestBody)

	item := Order{
		SellerId:   int(requestBody.SellerId),
		ReceiverId: int(user.ID),
		Item:       requestBody.Item,
		ItemAmount: requestBody.ItemAmount,
		Weight:     requestBody.Weight,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	database.Instance.Omit("Seller").Omit("Receiver").Create(&item)

	AddDetails(item)
}
