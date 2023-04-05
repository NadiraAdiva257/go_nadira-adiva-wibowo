package main

import (
	"middleware/config"
	"middleware/route"
)

func main() {
	config.Init()

	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
