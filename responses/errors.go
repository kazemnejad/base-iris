package responses

import (
	"yonje/baseframework/responses/code"
	"yonje/baseframework/valid"
)

type ErrorResponse struct {
	GeneralResponse
	Code int `json:"code"`
}

func NewError(code int) ErrorResponse {
	return ErrorResponse{
		GeneralResponse: GeneralResponse{"fail"},
		Code:            code,
	}
}

type ValidationErrorResponse struct {
	ErrorResponse
	Errors valid.Errors `json:"errors"`
}

func NewValidationError(errors valid.Errors) ValidationErrorResponse {
	return ValidationErrorResponse{
		ErrorResponse: NewError(code.BadInput),
		Errors:        errors,
	}
}
