package controller

import (
	"github.com/kataras/iris"
	"../models"
	"strconv"
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
	categoryID, _ := strconv.Atoi(ctx.FormValue("category_id"))

	goods := &models.Goods{Name:name, Description:description, Price:price, CreatedAt:time.Now(), CategoryID:categoryID}
	// user category 学习gorm之后再回来修改
	models.Db.Create(goods)

	var dict = map[string]int{"goods_id":goods.ID}
	ctx.JSON(dict)
}

func GetGoods(ctx iris.Context){
	category_id := ctx.FormValue("category_id")
	page := ctx.FormValue("page")
	page_size := ctx.FormValue("page_size")

	var goods []models.Goods

	// page从0开始
	pageint, _ := strconv.Atoi(page)

	page_sizeint, _ := strconv.Atoi(page_size)
	if category_id == "-1"{
		models.Db.Offset(pageint * page_sizeint).Limit(page_size).Order("id desc").Find(&goods)
	}else{
		models.Db.Where("category_id = ?", category_id).Offset(pageint * page_sizeint).Limit(page_size).Order("id desc").Find(&goods)
	}

	ctx.JSON(goods)
}

func GetAGoods(ctx iris.Context){
	goods_id := ctx.FormValue("goods_id")

	var goods []models.Goods
}
