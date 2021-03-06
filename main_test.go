package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgquijote/be-screening/controllers"
	"github.com/dgquijote/be-screening/database"
	"github.com/dgquijote/be-screening/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	connectionString := os.Getenv("DATABASE_URL")
	// Initialize Database
	database.MockConnect(connectionString)
	models.MigrateUsers()
	models.MigrateOrders()
	models.MigrateOrderDetails()
}

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func TestGenerateNoRequestHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/token", controllers.GenerateToken)

	req, _ := http.NewRequest("POST", "/api/token", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGenerateTokenInvalidUserHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/token", controllers.GenerateToken)

	user := controllers.TokenRequest{
		Email:    "someuser@email.com",
		Password: "000000000",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/token", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGenerateTokenInvalidPasswordHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/token", controllers.GenerateToken)

	user := controllers.TokenRequest{
		Email:    "test.user@email.com",
		Password: "000000000",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/token", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGenerateTokenHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/token", controllers.GenerateToken)

	user := controllers.TokenRequest{
		Email:    "test.user@email.com",
		Password: "123456789",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/token", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
