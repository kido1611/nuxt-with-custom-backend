package config

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "modernc.org/sqlite"
)

func NewDB(config *viper.Viper, log *logrus.Logger) *sql.DB {
	db, err := sql.Open(config.GetString("database.driver"), config.GetString("database.host"))
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
