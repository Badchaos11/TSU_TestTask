package main

import (
	"context"

	"github.com/Badchaos11/TSU_TestTask/config"
	"github.com/Badchaos11/TSU_TestTask/service"
)

func main() {
	ctx := context.Background()
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	app, err := service.NewService(ctx)
	if err != nil {
		panic(err)
	}

	app.Run()
}
