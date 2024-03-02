package main

import (
	"localdothost-context-menu/app"
	"log"
	"os"
)

func main() {
	app := app.Init()
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
