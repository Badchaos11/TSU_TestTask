package config

import (
	"os"

	"github.com/Badchaos11/TSU_TestTask/model"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() (*model.Config, error) {

	err := godotenv.Load("./configs/.env")
	if err != nil {
		logrus.Errorf("failed to load config: %v", err)
		return nil, err
	}

	return &model.Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}, nil
}
