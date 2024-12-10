package controllers

import (
	"abcShop/services/createProductService"
	"abcShop/services/getProductListService"
	"abcShop/services/getProductService"
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
}

func NewProductController(
	create createProductService.ICreateProductService,
	getlist getProductListService.IGetProductListService,
	get getProductService.IGetProductService,
) IProductController {
	return &productController{
		CreateProductService:  create,
		GetProductListService: getlist,
		GetProductService:     get,
	}
}

func (ctl *productController) Create(c *gin.Context) {
	var request createProductService.Request
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, map[string]string{"message": err.Error()})
		return
	}
	response, err := ctl.CreateProductService.Execute(request)
	if err != nil {
		c.JSON(400, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctl *productController) Find(c *gin.Context) {
	var request getProductListService.Request
	response, err := ctl.GetProductListService.Execute(request)
	if err != nil {
		c.JSON(400, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctl *productController) FindOne(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, map[string]string{"message": err.Error()})
		return
	}
	reqest := getProductService.Request{Id: id}

	response, err := ctl.GetProductService.Execute(reqest)
	if err != nil {
		c.JSON(400, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctl *productController) Update(c *gin.Context) {}

func (ctl *productController) Delete(c *gin.Context) {}
