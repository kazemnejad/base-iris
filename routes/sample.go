package routes

import (
	"github.com/kataras/iris"
	"yonje/baseframework/controllers"
)

func sampleControllerFactory() controllers.IController {
	return &controllers.SampleController{}
}

func getSampleController(ctx *iris.Context) *controllers.SampleController {
	return getControllerProvider(ctx).Provide("SampleController", sampleControllerFactory).(*controllers.SampleController)
}

func registerSampleRoutes(app *iris.Framework) {
	app.Get("/sample/:fullname", func(ctx *iris.Context) {
		getSampleController(ctx).Index(ctx, ctx.Param("fullname"))
	})
}
