package main

import (
	"log"
	"os"

	"github.com/dgquijote/be-screening/controllers"
	"github.com/dgquijote/be-screening/database"
	"github.com/dgquijote/be-screening/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("DATABASE_URL")
	// Initialize Database
	database.Connect(connectionString)
	database.Migrate()

	// Initialize Router
	router := initRouter()
	port := os.Getenv("PORT")
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
			// secured.GET("/:id", controllers.GetOrderById)
			secured.POST("/", controllers.CreateOrder())
		}
	}
	return router
}
