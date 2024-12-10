package controllers

import (
	"abcShop/services/updateUserPasswordService"
	"abcShop/services/updateUserService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Update(c *gin.Context)
	UpdatePassword(c *gin.Context)
}

type userController struct {
	UpdateUser         updateUserService.IUpdateUserService
	UpdateUserPassword updateUserPasswordService.IUpdateUserPasswordService
}

func NewUserController(
	updateUser updateUserService.IUpdateUserService,
	updateUserPassword updateUserPasswordService.IUpdateUserPasswordService,
) IUserController {
	return &userController{
		UpdateUser:         updateUser,
		UpdateUserPassword: updateUserPassword,
	}
}

func (ctl *userController) Update(c *gin.Context) {
	var request updateUserService.Request
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	response, err := ctl.UpdateUser.Execute(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctl *userController) UpdatePassword(c *gin.Context) {
	var request updateUserPasswordService.Request
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	response, err := ctl.UpdateUserPassword.Execute(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
