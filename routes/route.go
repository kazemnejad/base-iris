package routes

import (
	"github.com/kataras/iris"
	CONST "yonje/baseframework/constant"
	"yonje/baseframework/controllers"
)

func getControllerProvider(ctx *iris.Context) *controllers.ControllerProvider {
	return ctx.Get(CONST.ControlProv).(*controllers.ControllerProvider)
}

func RegisterRoutes(app *iris.Framework) {
	registerSampleRoutes(app)
}
