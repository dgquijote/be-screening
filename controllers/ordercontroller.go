package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgquijote/be-screening/models"

	"github.com/gin-gonic/gin"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var results = models.GetAll(c)
		c.IndentedJSON(http.StatusOK, results)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		user, record := models.GetUserByToken(tokenString)

		if record.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
			c.Abort()
			return
		}

		fmt.Println(user.ID)
		// requestBody := models.Order{}
		// c.BindJSON(&requestBody)

	}
}
