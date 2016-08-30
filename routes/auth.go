package routes

import (
	"github.com/kataras/iris"
	"yonje/baseframework/controllers"
	"yonje/baseframework/helpers"
)

func authControllerFactory() controllers.IController {
	return &controllers.AuthController{}
}

func getAuthController(ctx *iris.Context) *controllers.AuthController {
	return getControllerProvider(ctx).Provide("AuthController", authControllerFactory).(*controllers.AuthController)
}

func registerAuthRoutes(app *iris.Framework) {
	app.Post("/auth/register", func(ctx *iris.Context) {
		getAuthController(ctx).Register(helpers.NewRequestWrapper(ctx))
	})

	app.Post("/auth/login", func(ctx *iris.Context) {
		getAuthController(ctx).Login(helpers.NewRequestWrapper(ctx))
	})
}
