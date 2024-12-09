package controllers

import (
	"abcShop/services/createProductService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
	FindOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type productController struct {
	CreateProductService createProductService.ICreateProductService
}

func NewProductController(c createProductService.ICreateProductService) IProductController {
	return &productController{
		CreateProductService: c,
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

func (ctl *productController) Find(c *gin.Context) {}

func (ctl *productController) FindOne(c *gin.Context) {}

func (ctl *productController) Update(c *gin.Context) {}

func (ctl *productController) Delete(c *gin.Context) {}