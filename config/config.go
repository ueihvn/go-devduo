package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	PostgresUser     string
	PostgresPassword string
	PostgrePort      string
	PostgresHost     string
	PostgresDb       string
}

func LoadEnv() error {
	err := godotenv.Load(".env.development")
	if err != nil {
		return err
	}
	return nil
}

func NewConfig() *DbConfig {
	err := LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	return &DbConfig{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgrePort:      os.Getenv("POSTGRES_PORT"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresDb:       os.Getenv("POSTGRES_DB"),
	}
}
