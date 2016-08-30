package controllers

import (
	"yonje/baseframework/helpers"
	"yonje/baseframework/models"
)

type SampleController struct{}

func (self *SampleController) Init() {}

func (self *SampleController) DeInit() {}

func (self *SampleController) Index(request *helpers.RequestWrapper, param string) {
	request.Context.JSON(200, struct {
		Text string `json:"text"`
	}{"This is test index: " + param})
}

func (self *SampleController) Protected(request *helpers.RequestWrapper) {
	ctx := request.Context

	request.Context.JSON(200, struct {
		Text string       `json:"text"`
		User *models.User `json:"user"`
	}{"Protected", auth().User(ctx)})
}
