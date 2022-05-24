package main

import (
	"delivery-app/controllers"
	"delivery-app/database"
	"delivery-app/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root:@tcp(localhost:3306)/delivery_app?parseTime=true")
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
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
