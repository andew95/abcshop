package routes

import (
	"abcShop/controllers"

	"github.com/gin-gonic/gin"
)

func addProductRouter(rg *gin.RouterGroup, ctl controllers.IProductController) {
	r := rg.Group("product")
	{
		r.POST("", ctl.Create)
		r.GET("", ctl.Find)
		r.GET("/:id", ctl.FindOne)
		r.PUT("", ctl.Update)
		r.DELETE("", ctl.Delete)
	}
}
