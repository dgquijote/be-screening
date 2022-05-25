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

	dburl := os.Getenv("DATABASE_URL")
	dbuser := os.Getenv("DATABASE_USER")
	dbpass := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	dbport := os.Getenv("DATABASE_PORT")

	connectionString := "host=" + dburl + " user=" + dbuser + " password=" + dbpass + " dbname=" + dbname + " port=" + dbport + " sslmode=require"
	// Initialize Database
	database.Connect(connectionString)
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run("localhost:8000")
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
