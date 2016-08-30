package controllers

import (
	"yonje/baseframework/db"
	"yonje/baseframework/db/tables"
	"yonje/baseframework/helpers"
	"yonje/baseframework/models"
	"yonje/baseframework/responses"
	"yonje/baseframework/responses/code"
	"yonje/baseframework/valid"
)

type AuthController struct {
	*valid.Validator
}

func (self *AuthController) Init() {
	self.Validator = valid.NewValidator()
}

func (self *AuthController) DeInit() {}

func (self *AuthController) Register(request *helpers.RequestWrapper) {
	ctx := request.Context

	if errs, ok := self.Validate(request, self.getValidationRules()); !ok {
		ctx.JSON(400, responses.NewValidationError(errs))
		return
	}

	user := models.User{
		Email:    request.PostValue("email"),
		Password: models.GeneratePasswordHash(request.PostValue("password")),
		Name:     request.PostValue("name"),
	}

	result, err := db.Session().InsertInto(tables.User).Columns("email", "password", "name").Record(user).Exec()
	if err != nil {
		ctx.JSON(400, responses.NewValidationError(valid.Errors{
			"email": valid.DefaultErrors["unique"],
		}))
		return
	}

	user.Id, _ = result.LastInsertId()

	ctx.JSON(200, responses.NewAuth(user.Id, user.GenerateJwtToken()))
}

func (self *AuthController) Login(request *helpers.RequestWrapper) {
	ctx := request.Context

	rules := valid.Rules{
		"email":    "required",
		"password": "required",
	}
	if errs, ok := self.Validate(request, rules); !ok {
		ctx.JSON(400, responses.NewValidationError(errs))
		return
	}

	user := models.User{}
	db.Session().Select("id", "password").
		From(tables.User).
		Where("email = ?", request.PostValue("email")).
		LoadStruct(&user)

	if user.Password != "" && models.CheckHashWithPassword(user.Password, request.PostValue("password")) {
		ctx.JSON(200, responses.NewAuth(user.Id, user.GenerateJwtToken()))
		return
	}

	ctx.JSON(403, responses.NewError(code.UnmatchedUserPass))
}

func (self *AuthController) getValidationRules() valid.Rules {
	return valid.Rules{
		"email":    "required,email",
		"password": "required",
		"name":     "required",
	}
}
