package main

import (
	"github.com/AppleBoiy/pydumper/go/crawler/app"
	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/config"
	"github.com/AppleBoiy/pydumper/go/crawler/app/model"
)

func main() {
	// load config.ini
	config.LoadConfig()

	// connect to database
	model.MD.Initialize()

	// start Application
	app.Start()
}
