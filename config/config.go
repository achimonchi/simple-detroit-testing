package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig(filenames ...string) error {
	err := godotenv.Load(filenames...)
	return err
}

func GetString(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}
func GetUint8(key string, fallback uint8) uint8 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	valUint8, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return uint8(valUint8)
}
