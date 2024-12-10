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
	// register product port
	product := controllers.NewProductController(
		service.CreateProduct,
		service.GetProductList,
		service.GetProduct,
		service.UpdateProduct,
		service.DeleteProduct,
	)
	return &SetupController{
		RegisterController: register,
		LoginController:    login,
		ProductController:  product,
	}
}
