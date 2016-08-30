package authmiddleware

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"strconv"
	"yonje/baseframework/config"
	"yonje/baseframework/db"
	"yonje/baseframework/db/tables"
	"yonje/baseframework/models"
	"yonje/baseframework/responses"
	"yonje/baseframework/responses/code"
	"yonje/baseframework/shared"
)

var instance *Middleware = nil

type Middleware struct {
	*jwtmiddleware.Middleware
}

func (self *Middleware) Serve(ctx *iris.Context) {
	err := self.CheckJWT(ctx)

	if err != nil {
		ctx.Log(err)
		ctx.JSON(403, responses.NewError(code.PermissionDenied))
		return
	}

	ctx.Set(shared.LoadedUser, nil)
	ctx.Next()
}

func (self *Middleware) User(ctx *iris.Context) *models.User {
	loadedUser := ctx.Get(shared.LoadedUser)
	if loadedUser == nil {
		userId, _ := strconv.ParseInt(self.Get(ctx).Claims.(jwt.MapClaims)["userId"].(string), 10, 64)

		user := models.User{}
		db.Session().Select("*").From(tables.User).Where("id = ?", userId).LoadStruct(&user)

		ctx.Set(shared.LoadedUser, &user)
		loadedUser = &user
	}

	return loadedUser.(*models.User)
}

func newAuth() *Middleware {
	return &Middleware{
		Middleware: jwtmiddleware.New(jwtmiddleware.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return config.SigningSecret, nil
			},
			SigningMethod: jwt.SigningMethodHS256,
			ContextKey:    shared.JwtToken,
		}),
	}
}

func Get() *Middleware {
	if instance == nil {
		instance = newAuth()
	}

	return instance
}
