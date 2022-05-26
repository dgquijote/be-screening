package models

import (
	"os"
	"testing"

	"github.com/dgquijote/be-screening/database"
	"github.com/stretchr/testify/assert"
)

func init() {
	connectionString := os.Getenv("DATABASE_URL")
	database.MockConnect(connectionString)
}

func TestGetOrderTrackingDetails(t *testing.T) {
	assert.Equal(t, GetOrderTrackingDetails(1)[0].Id, uint(1))
}
