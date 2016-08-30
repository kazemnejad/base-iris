package valid

import (
	"gopkg.in/go-playground/validator.v8"
	"yonje/baseframework/helpers"
)

type Rules map[string]string
type Errors map[string]Error

type Validator struct {
	engine *validator.Validate
}

func (self *Validator) Validate(request *helpers.RequestWrapper, rules Rules) (Errors, bool) {
	isOk := true
	errs := make(Errors)

	for fieldName, rule := range rules {
		err := self.engine.Field(self.getValue(request, fieldName), rule)
		if err != nil {
			isOk = false
			for _, value := range err.(validator.ValidationErrors) {
				errs[fieldName] = DefaultErrors[value.Tag]
				break
			}
		}
	}

	return errs, isOk
}

func (self *Validator) getValue(request *helpers.RequestWrapper, name string) interface{} {
	values := request.PostValues(name)
	switch length := len(values); length {
	case 0:
		return nil
	case 1:
		return values[0]
	default:
		return values
	}
}

func NewValidator() *Validator {
	return &Validator{
		engine: validator.New(&validator.Config{
			TagName: "valid",
		}),
	}
}
