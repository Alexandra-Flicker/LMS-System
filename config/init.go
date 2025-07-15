package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️ .env не найден, читаю переменные окружения напрямую")
	}

	cfg := &Config{}

	return cfg
}
