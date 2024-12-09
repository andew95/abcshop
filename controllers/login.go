package controllers

import (
	"abcShop/Dtos/loginDto"
	"abcShop/services/loginService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	Login(c *gin.Context)
}

type loginController struct {
	LoginService loginService.ILoginService
}

func NewLoginController(loginService loginService.ILoginService) ILoginController {
	return &loginController{
		LoginService: loginService,
	}
}

func (ctl *loginController) Login(c *gin.Context) {
	var request loginDto.Request

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	response, err := ctl.LoginService.Login(request.ToLoginServiceRequest())
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
