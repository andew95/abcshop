package routes

import (
	"abcShop/controllers"

	"github.com/gin-gonic/gin"
)

func addUserRouter(rg *gin.RouterGroup, ctl controllers.IUserController) {
	r := rg.Group("user")
	{
		r.PUT("", ctl.Update)
		r.PATCH("", ctl.UpdatePassword)
	}
}
