package routes

import (
	"github.com/kataras/iris"
	"yonje/baseframework/controllers"
	"yonje/baseframework/helpers"
	"yonje/baseframework/middlewares/authmiddleware"
	"yonje/baseframework/shared"
)

func auth(ctx *iris.Context) {
	authmiddleware.Get().Serve(ctx)
}

func request(ctx *iris.Context) *helpers.RequestWrapper {
	return helpers.NewRequestWrapper(ctx)
}

func getControllerProvider(ctx *iris.Context) *controllers.ControllerProvider {
	return ctx.Get(shared.ControlProv).(*controllers.ControllerProvider)
}

func RegisterRoutes(app *iris.Framework) {
	registerSampleRoutes(app)
	registerAuthRoutes(app)
}
