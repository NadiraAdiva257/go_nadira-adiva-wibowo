package main

import (
	"clean-architecture/cmd/app"
	"log"
)

func main() {
	log.Println("Starting application...")
	route := app.StartApp()

	route.Start(":8000")
}
