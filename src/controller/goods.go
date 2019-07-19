package controller

import (
	"github.com/kataras/iris"
	"../models"
	"time"
)
func GetSubCategoryGoods(ctx iris.Context){
	categoryId := ctx.FormValue("id")
	var Goods_ []models.Goods

	models.Db.Find(&Goods_, "category_id = ?", categoryId)

	ctx.JSON(Goods_)
}

func AddGoods(ctx iris.Context){
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	price := ctx.FormValue("price")

	// user category 学习gorm之后再回来修改
	models.Db.Create(&models.Goods{Name:name, Description:description, Price:price, CreatedAt:time.Now()})
}