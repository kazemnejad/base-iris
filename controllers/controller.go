package controllers

type Controller interface {
	Init()
	DeInit()
}

type ControllerFactory func() Controller

type ControllerProvider struct {
	cache map[string]Controller
}

func (self *ControllerProvider) Provide(name string, factory ControllerFactory) Controller {
	var controller Controller
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
		cache: make(map[string]Controller),
	}
}
