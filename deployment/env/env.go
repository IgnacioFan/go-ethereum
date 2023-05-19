package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func EthClientURL() (string, error) {
	if err := LoadEnv(); err != nil {
		return "", err
	}
	return os.Getenv("ETH_CLIENT_URL"), nil
}

func DSN() (string, error) {
	if err := LoadEnv(); err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	), nil
}
