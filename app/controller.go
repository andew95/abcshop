package app

import "abcShop/controllers"

type SetupController struct {
	RegisterController controllers.IRegisterController
	LoginController    controllers.ILoginController
	ProductController  controllers.IProductController
}

func NewSetupController(service *SetupService) *SetupController {
	register := controllers.NewRegisterController(service.Register)
	login := controllers.NewLoginController(service.Login)
	product := controllers.NewProductController(service.CreateProduct, service.GetProductList)
	return &SetupController{
		RegisterController: register,
		LoginController:    login,
		ProductController:  product,
	}
}
