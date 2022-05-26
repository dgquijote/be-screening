package controllers

import (
	"net/http"

	"github.com/dgquijote/be-screening/models"

	"github.com/gin-gonic/gin"
)

type OrderReponse struct {
	OrderDetails    models.Order
	TrackingDetails []models.OrderDetails
}

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var results = models.GetAll(c)
		c.IndentedJSON(http.StatusOK, results)
	}
}

func GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		o := models.GetById(c)
		if o.Id == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "order not found"})
			return
		}
		d := models.GetOrderTrackingDetails(o.Id)
		var result = OrderReponse{
			OrderDetails:    o,
			TrackingDetails: d,
		}

		c.IndentedJSON(http.StatusOK, result)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		models.Add(c)
		c.JSON(http.StatusCreated, gin.H{"message": "Order has been checked out."})
	}
}
