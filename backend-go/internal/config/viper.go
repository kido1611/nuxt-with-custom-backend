package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable load .env file")
	}

	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
	config.AddConfigPath("../../")

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	err = config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed when reading config: %w", err))
	}

	return config
}
