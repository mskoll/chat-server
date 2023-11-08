package main

import (
	"chat-server/internal/app"
	"log"
)

func main() {

	app, err := app.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	app.Run()

}
