package app

import (
	"abcShop/services/createProductService"
	"abcShop/services/getProductListService"
	"abcShop/services/loginService"
	"abcShop/services/registerService"
)

type SetupService struct {
	Register       registerService.IRegisterService
	Login          loginService.ILoginService
	CreateProduct  createProductService.ICreateProductService
	GetProductList getProductListService.IGetProductListService
}

func NewSetupService(repo *SetupRepository) *SetupService {
	register := registerService.NewRegisterService(repo.UserRepo)
	login := loginService.NewLoginService(repo.UserRepo)

	// product service
	createProduct := createProductService.NewCreateProductService(repo.ProductRepo)
	getProductList := getProductListService.NewGetProductListService(repo.ProductRepo)

	return &SetupService{
		Register:       register,
		Login:          login,
		CreateProduct:  createProduct,
		GetProductList: getProductList,
	}
}
