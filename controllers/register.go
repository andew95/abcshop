package controllers

import (
	"abcShop/Dtos/registerDto"
	"abcShop/services/registerService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRegisterController interface {
	Register(c *gin.Context)
}

type registerController struct {
	RegisterService registerService.IRegisterService
}

func NewRegisterController(registerService registerService.IRegisterService) IRegisterController {
	return &registerController{
		RegisterService: registerService,
	}
}

func (ctl *registerController) Register(c *gin.Context) {
	var request registerDto.Request
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	response, err := ctl.RegisterService.Execute(request)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
