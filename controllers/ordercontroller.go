package controllers

import (
	"net/http"

	"github.com/dgquijote/be-screening/models"

	"github.com/gin-gonic/gin"
)

func GetOrders(context *gin.Context) {
	var orders models.Order
	context.IndentedJSON(http.StatusOK, orders)
}
