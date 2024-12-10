package app

import "abcShop/controllers"

type SetupController struct {
	RegisterController controllers.IRegisterController
	LoginController    controllers.ILoginController
	ProductController  controllers.IProductController
	UserController     controllers.IUserController
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
	// register user port
	user := controllers.NewUserController(service.UpdateUser, service.UpdateUserPassword)
	return &SetupController{
		RegisterController: register,
		LoginController:    login,
		ProductController:  product,
		UserController:     user,
	}
}
