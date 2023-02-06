package main

import "github.com/Badchaos11/TSU_TestTask/service"

func main() {
	app, err := service.NewService()
	if err != nil {
		panic(err)
	}

	app.Run()
}
