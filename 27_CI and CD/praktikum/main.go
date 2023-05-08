package main

import (
	"orm-mvc-eksplorasi/config"
	"orm-mvc-eksplorasi/route"
)

func main() {
	config.Init()

	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
