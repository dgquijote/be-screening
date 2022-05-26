package main

import (
	"os"

	"github.com/dgquijote/be-screening/controllers"
	"github.com/dgquijote/be-screening/database"
	"github.com/dgquijote/be-screening/middlewares"
	"github.com/dgquijote/be-screening/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		connectionString = "host=localhost user=postgres password= dbname=delivery_app port=5432"
	}
	database.Connect(connectionString)
	models.MigrateUsers()
	models.MigrateOrders()
	models.MigrateOrderDetails()

	// Initialize Router
	router := initRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router.Run(":" + port)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		secured := api.Group("/order").Use(middlewares.Auth())
		{
			secured.GET("/", controllers.GetOrders())
			secured.GET("/:id", controllers.GetOrderById())
			secured.POST("/", controllers.CreateOrder())
		}
	}
	return router
}
