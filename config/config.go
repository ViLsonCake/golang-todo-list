package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type WebConfig struct {
	Port      string
	JWTSecret string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetWebConfig() WebConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return WebConfig{
		Port:      os.Getenv("API_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

func GetPostgresConfig() PostgresConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DATABASE"),
	}
}

func GetPostgresUrl() string {
	postgresConfig := GetPostgresConfig()
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresConfig.User, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database)
}
