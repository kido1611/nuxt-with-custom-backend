package main

import (
	"fmt"
	"kido1611/notes-backend-go/internal/config"
	"kido1611/notes-backend-go/internal/delivery/http"
)

func main() {
	viper := config.NewViper()
	validate := config.NewValidator()
	log := config.NewLogrus(viper)
	db := config.NewDB(viper, log)
	app := config.NewFiber()

	router := http.NewRouter(db, app, validate, viper, log)
	router.Setup()

	err := app.Listen(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
