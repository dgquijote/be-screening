package controllers

import (
	"delivery-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(context *gin.Context) {
	var orders models.Order
	context.IndentedJSON(http.StatusOK, orders)
}
