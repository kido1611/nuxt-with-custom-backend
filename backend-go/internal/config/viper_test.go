package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	config := NewViper()

	assert.Equal(t, "db.sqlite", config.GetString("database.host"))
	assert.Equal(t, "sqlite", config.GetString("database.driver"))
}

func TestEnvOverride(t *testing.T) {
	os.Setenv("DATABASE_HOST", "test-db.sqlite")
	config := NewViper()

	assert.NotEqual(t, "db.sqlite", config.GetString("database.host"))
	assert.Equal(t, "test-db.sqlite", config.GetString("database.host"))
}
