package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	config := NewViper()

	assert.Equal(t, "sqlite://db.sqlite", config.GetString("database.url"))
}

func TestEnvOverride(t *testing.T) {
	os.Setenv("DATABASE_URL", "sqlite://test-db.sqlite")
	config := NewViper()

	assert.NotEqual(t, "sqlite://db.sqlite", config.GetString("database.url"))
	assert.Equal(t, "sqlite://test-db.sqlite", config.GetString("database.url"))
}
