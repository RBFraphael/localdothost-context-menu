package main

import (
	"localdothost-context-menu/app"
	"os"
)

func main() {
	app := app.Init()
	app.Run(os.Args)
}