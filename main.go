package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"./src/controller"
	"./src/models"
)

func main() {

	models.Db.AutoMigrate(&models.User{}, &models.Category{}, &models.Goods{})

	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	app.Post("/bind", controller.Bind)
	app.Post("/getParentCategory", controller.GetParentCategory)
	app.Post("/getSubCategory", controller.GetSubCategory)
	app.Post("/getSubCategoryGoods", controller.GetSubCategoryGoods)
	app.Post("/addGoods", controller.AddGoods)

	app.Run(iris.Addr(":80"), iris.WithoutServerError(iris.ErrServerClosed))
}
