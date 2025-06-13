package main

import (
	"fmt"
	"kido1611/notes-backend-go/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	viper := config.NewViper()
	// https://github.com/golang-migrate/migrate?tab=readme-ov-file#use-in-your-go-project
	// log := config.NewLogrus(viper)
	// db := config.NewDB(viper, log)
	schema := viper.GetString("database.driver") + "://" + viper.GetString("database.host")

	m, err := migrate.New("file://internal/db/migrations", schema)
	if err != nil {
		panic(fmt.Errorf("failed when reading migrations: %w", err))
	}

	_ = m.Up()
}
