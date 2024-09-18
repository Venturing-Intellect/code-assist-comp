package configs

import (
	"fmt"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBSSLMode  string
	Port       string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSL"),
		Port:       os.Getenv("PORT"),
	}

	if config.DBUser == "" || config.DBPassword == "" || config.DBHost == "" || config.DBPort == "" || config.DBName == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	if config.DBSSLMode == "" {
		config.DBSSLMode = "disable"
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	return config, nil
}

func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}
