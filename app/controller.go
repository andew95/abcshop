package app

import "abcShop/controllers"

type SetupController struct {
	RegisterController controllers.IRegisterController
	LoginController    controllers.ILoginController
}

func NewSetupController(service *SetupService) *SetupController {
	register := controllers.NewRegisterController(service.Register)
	login := controllers.NewLoginController(service.Login)
	return &SetupController{
		RegisterController: register,
		LoginController:    login,
	}
}
