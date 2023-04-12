package main

import (
	"docker/config"
	"docker/route"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.Init()

	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
