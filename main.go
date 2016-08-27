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
	app.ListenTo(config.GetServerConfig())

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

	if db.Connection != nil {
		db.Connection.Close()
	}
}

func makeApp() *iris.Framework {
	app := iris.New(config.GetIrisConfig())

	middlewares.RegisterGlobalMiddlewares(app)
	routes.RegisterRoutes(app)

	return app
}
