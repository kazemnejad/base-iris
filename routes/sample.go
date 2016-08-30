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
	app.Get("/sample/index/:fullname", func(ctx *iris.Context) {
		getSampleController(ctx).Index(request(ctx), ctx.Param("fullname"))
	})

	app.Get("/sample/protected", auth, func(ctx *iris.Context) {
		getSampleController(ctx).Protected(request(ctx))
	})
}
