package app

import (
	"abcShop/services/createProductService"
	"abcShop/services/deleteProductService"
	"abcShop/services/getProductListService"
	"abcShop/services/getProductService"
	"abcShop/services/loginService"
	"abcShop/services/registerService"
	"abcShop/services/updateProductService"
)

type SetupService struct {
	Register       registerService.IRegisterService
	Login          loginService.ILoginService
	CreateProduct  createProductService.ICreateProductService
	GetProductList getProductListService.IGetProductListService
	GetProduct     getProductService.IGetProductService
	UpdateProduct  updateProductService.IUpdateProductService
	DeleteProduct  deleteProductService.IDeleteProductService
}

func NewSetupService(repo *SetupRepository) *SetupService {
	register := registerService.NewRegisterService(repo.UserRepo)
	login := loginService.NewLoginService(repo.UserRepo)

	// product service
	createProduct := createProductService.NewCreateProductService(repo.ProductRepo)
	getProductList := getProductListService.NewGetProductListService(repo.ProductRepo)
	getProduct := getProductService.NewCreateProductService(repo.ProductRepo)
	updateProduct := updateProductService.NewUpdateProductService(repo.ProductRepo)
	deleteProduct := deleteProductService.NewDeleteProductService(repo.ProductRepo)

	return &SetupService{
		Register:       register,
		Login:          login,
		CreateProduct:  createProduct,
		GetProductList: getProductList,
		GetProduct:     getProduct,
		UpdateProduct:  updateProduct,
		DeleteProduct:  deleteProduct,
	}
}
