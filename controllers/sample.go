package controllers

import "github.com/kataras/iris"

type SampleController struct{}

func (self *SampleController) Init() {}

func (self *SampleController) DeInit() {}

func (self *SampleController) Index(app *iris.Context, param string) {
	app.JSON(200, struct {
		Text string `json:"text"`
	}{"This is test index: " + param})
}
