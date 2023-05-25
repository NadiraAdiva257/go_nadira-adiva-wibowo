package main

import (
	"code-competence/config"
	"code-competence/route"
)

func main() {
	config.Init()

	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
