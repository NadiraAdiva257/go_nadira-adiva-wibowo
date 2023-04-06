package main

import (
	"orm-mvc/prioritas1/config"
	"orm-mvc/prioritas1/route"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.Init()

	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
