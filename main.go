package main

import (
	"fmt"
	"github.com/kataras/iris"
	"yonje/baseframework/config"
	"yonje/baseframework/db"
	"yonje/baseframework/middlewares"
	"yonje/baseframework/routes"
)

func main() {
	if ok := initialize(); !ok {
		fmt.Println("Unable to initialize app")
		return
	}

	app := makeApp()
	err := app.ListenTo(config.GetServerConfig())

	fmt.Println(err)

	deInitialize()
}

func initialize() bool {
	if err := db.OpenGlobalConnection(); err != nil {
		return false
	}

	return true
}

func deInitialize() {
	fmt.Println("Deinitializing...")

	if db.Session() != nil {
		db.Session().Close()
	}
}

func makeApp() *iris.Framework {
	app := iris.New(config.GetIrisConfig())

	middlewares.RegisterGlobalMiddlewares(app)
	routes.RegisterRoutes(app)

	return app
}
