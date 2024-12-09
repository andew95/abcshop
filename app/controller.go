package app

import "abcShop/controllers"

type SetupController struct {
	RegisterController controllers.IRegisterController
}

func NewSetupController(service *SetupService) *SetupController {
	register := controllers.NewRegisterController(service.Register)
	return &SetupController{
		RegisterController: register,
	}
}
