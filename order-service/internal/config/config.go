package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Settings struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	Ssl        string

	HttpPort           string
	TokenPrefix        string
	TokenLength        int
	PaymentServiceURL  string
}

func (s *Settings) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		s.DbUser, s.DbPassword, s.DbHost, s.DbPort, s.DbName, s.Ssl)
}

func LoadConfig() (*Settings, error) {
	godotenv.Load(".env")

	return &Settings{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		Ssl:        os.Getenv("DB_SSL"),

		HttpPort:           os.Getenv("HTTP_PORT"),
		TokenPrefix:        os.Getenv("TOKEN_PREFIX"),
		TokenLength:        getEnvAsInt("TOKEN_LENGTH", 32),
		PaymentServiceURL:  os.Getenv("PAYMENT_SERVICE_URL"),
	}, nil
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if valueStr == "" {
		return defaultVal
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultVal
	}

	return value
}
