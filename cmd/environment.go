package main

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func validateEnvironments() error {
	if strings.TrimSpace(os.Getenv("SERVER_PORT")) == "" {
		return errors.New("the SERVER_PORT env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS")) == "" {
		return errors.New("the ALLOWED_ORIGINS env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_METHODS")) == "" {
		return errors.New("the ALLOWED_METHODS env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("IMAGES_DIR")) == "" {
		return errors.New("the IMAGES_DIR env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("JWT_SECRET_KEY")) == "" {
		return errors.New("the JWT_SECRET_KEY env is mandatory")
	}

	// Database
	if strings.TrimSpace(os.Getenv("DB_USER")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_PASSWORD")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_HOST")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_PORT")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_NAME")) == "" {
		return errors.New("the env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("DB_SSL_MODE")) == "" {
		return errors.New("the env is mandatory")
	}

	// PayPal envs.
	if strings.TrimSpace(os.Getenv("WEBHOOK_ID")) == "" {
		return errors.New("the WEBHOOK_ID env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("VALIDATION_URL")) == "" {
		return errors.New("the VALIDATION_URL env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("CLIENT_ID")) == "" {
		return errors.New("the CLIENT_ID env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("SECRET_ID")) == "" {
		return errors.New("the SECRET_ID env is mandatory")
	}

	return nil
}
