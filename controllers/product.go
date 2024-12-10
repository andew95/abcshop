package controllers

import (
	"abcShop/services/createProductService"
	"abcShop/services/deleteProductService"
	"abcShop/services/getProductListService"
	"abcShop/services/getProductService"
	"abcShop/services/updateProductService"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IProductController interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
	FindOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type productController struct {
	CreateProductService  createProductService.ICreateProductService
	GetProductListService getProductListService.IGetProductListService
	GetProductService     getProductService.IGetProductService
	UpdateProductService  updateProductService.IUpdateProductService
	DeleteProductService  deleteProductService.IDeleteProductService
}

func NewProductController(
	create createProductService.ICreateProductService,
	getlist getProductListService.IGetProductListService,
	get getProductService.IGetProductService,
	update updateProductService.IUpdateProductService,
	deleteProduct deleteProductService.IDeleteProductService,
) IProductController {
	return &productController{
		CreateProductService:  create,
		GetProductListService: getlist,
		GetProductService:     get,
		UpdateProductService:  update,
		DeleteProductService:  deleteProduct,
	}
}

func (ctl *productController) Create(c *gin.Context) {
	var request createProductService.Request
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	response, err := ctl.CreateProductService.Execute(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctl *productController) Find(c *gin.Context) {
	var request getProductListService.Request
	response, err := ctl.GetProductListService.Execute(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctl *productController) FindOne(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	response, err := ctl.GetProductService.Execute(getProductService.Request{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctl *productController) Update(c *gin.Context) {
	var reqest updateProductService.Request
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	err = c.BindJSON(&reqest)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	reqest.Id = id
	response, err := ctl.UpdateProductService.Execute(reqest)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctl *productController) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	err = ctl.DeleteProductService.Execute(deleteProductService.Request{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "deleted product"})
}
