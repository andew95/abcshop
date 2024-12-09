package main

import (
	"abcShop/app"
	"abcShop/configs"
	"abcShop/db"
	"abcShop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := configs.GetConfig()
	db := db.NewSqlConnection(conf.DbConfig.User, conf.DbConfig.Password, conf.DbConfig.DbName)
	// if db != nil {
	// 	fmt.Println("success")
	// }

	repo := app.NewSetupRepository(db)
	service := app.NewSetupService(repo)
	controller := app.NewSetupController(service)

	r := gin.Default()

	routes.SetupRoute(r, controller)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
