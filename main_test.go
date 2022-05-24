package main

import (
	"bytes"
	"delivery-app/controllers"
	"delivery-app/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	database.MockConnect("root:@tcp(localhost:3306)/delivery_app?parseTime=true")
	database.Migrate()
}

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func TestGenerateNoRequestHandler(t *testing.T) {
	r := SetUpRouter()
	api := r.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
	}

	req, _ := http.NewRequest("POST", "/api/token", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGenerateTokenInvalidUserHandler(t *testing.T) {
	r := SetUpRouter()
	api := r.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
	}

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
	api := r.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
	}

	user := controllers.TokenRequest{
		Email:    "seller@email.com",
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
	api := r.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
	}

	user := controllers.TokenRequest{
		Email:    "seller@email.com",
		Password: "123465789",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/token", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserNotAuthenticatedHandler(t *testing.T) {
}
