package config

import (
	"os"

	"github.com/Badchaos11/TSU_TestTask/model"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() (*model.Config, error) {
	d := os.Getenv("DEPLOY")
	f := "./configs/.env"
	if d == "docker" {
		f = "./configs/docker.env"
	}

	err := godotenv.Load(f)
	if err != nil {
		logrus.Errorf("failed to load config: %v", err)
		return nil, err
	}

	return &model.Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		CacheUrl:   os.Getenv("CACHE_URL"),
	}, nil
}
