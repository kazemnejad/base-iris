package middlewares

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
	CONST "yonje/baseframework/constant"
	"yonje/baseframework/controllers"
	"yonje/baseframework/db"
)

func AddControllerProvider(ctx *iris.Context) {
	ctx.Set(CONST.ControlProv, controllers.NewControllerProvider())
	ctx.Next()
}

func AddDatabaseSession(ctx *iris.Context) {
	if db.Connection == nil {
		ctx.JSON(500, struct {
			Message string `json:"message"`
		}{"no open database connection"})
		return
	}

	ctx.Set(CONST.DbSession, db.Connection.NewSession(nil))
	ctx.Next()
}

func RegisterGlobalMiddlewares(app *iris.Framework) {
	// use default Iris logger to log each request
	app.Use(logger.New(app.Logger))

	// add dbr session object per request
	app.UseFunc(AddDatabaseSession)

	// add controllers cache to each request
	app.UseFunc(AddControllerProvider)
}
