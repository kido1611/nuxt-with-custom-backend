package config

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewDB(config *viper.Viper, log *logrus.Logger) *sql.DB {
	db, err := sql.Open(config.GetString("database.driver"), config.GetString("database.url"))
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	defer db.Close()

	return db
}
