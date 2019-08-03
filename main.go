package main

import (
	"./src/controller"
	"./src/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	models.Db.AutoMigrate(&models.User{}, &models.Category{}, &models.Goods{}, &models.Image{})

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
	app.Post("/addImage", controller.AddImage)
	app.Post("/getGoods", controller.GetGoods)
	// to start a new server listening at :80 and redirects
	// to the secure address, then:
	//target, _ := url.Parse("https://127.0.0.1:443")
	//go host.NewProxy("127.0.0.1:80", target).ListenAndServe()
	//
	//// start the server (HTTPS) on port 443, this is a blocking func,
	//// you can use iris.AutoTLS for letsencrypt if you want so.
	app.Run(iris.TLS("0.0.0.0:443", "neumark.pem", "neumark.key"))

	// app.Run(iris.Addr(":80"), iris.WithoutServerError(iris.ErrServerClosed))
}
