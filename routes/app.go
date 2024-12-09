package routes

import (
	"abcShop/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, ctl *app.SetupController) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("v1")
	{
		v1.POST("register", ctl.RegisterController.Register)
		v1.POST("login", ctl.LoginController.Login)
	}
}
