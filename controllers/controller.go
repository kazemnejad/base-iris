package controllers

import "yonje/baseframework/middlewares/authmiddleware"

type IController interface {
	Init()
	DeInit()
}

type ControllerFactory func() IController

type ControllerProvider struct {
	cache map[string]IController
}

func (self *ControllerProvider) Provide(name string, factory ControllerFactory) IController {
	var controller IController
	if val, ok := self.cache[name]; ok {
		controller = val
	} else {
		controller = factory()
		controller.Init()

		self.cache[name] = controller
	}

	return controller
}

func NewControllerProvider() *ControllerProvider {
	return &ControllerProvider{
		cache: make(map[string]IController),
	}
}

func auth() *authmiddleware.Middleware {
	return authmiddleware.Get()
}
