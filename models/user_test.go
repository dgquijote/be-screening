package models

import (
	"os"
	"testing"

	"github.com/dgquijote/be-screening/auth"
	"github.com/dgquijote/be-screening/database"
	"github.com/stretchr/testify/assert"
)

func init() {
	connectionString := os.Getenv("DATABASE_URL")
	database.MockConnect(connectionString)
}

func TestGetUserByToken(t *testing.T) {
	token, _ := auth.GenerateJWT("test.user@email.com", "test.user")
	user, _ := GetUserByToken(token)
	assert.Equal(t, user.Username, "test.user")
}
