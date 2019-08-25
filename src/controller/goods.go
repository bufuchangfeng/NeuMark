package controller

import (
	"../models"
	"github.com/kataras/iris"
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
	userID, _ := strconv.Atoi(ctx.FormValue("user_id"))

	categoryID, _ := strconv.Atoi(ctx.FormValue("category_id"))

	goods := &models.Goods{Name:name, Description:description, Price:price, CreatedAt:time.Now(), CategoryID:categoryID, UserID:userID}
	// user category 学习gorm之后再回来修改
	models.Db.Create(goods)

	sell := &models.Sell{UserID:userID, GoodsID:goods.ID, CreatedAt:time.Now()}
	models.Db.Create(sell)

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
	}else {
		models.Db.Where("category_id = ?", category_id).Offset(pageint * page_sizeint).Limit(page_size).Order("id desc").Find(&goods)
	}

	for i := 0; i < len(goods); i++ {
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Comments)
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Collects)
	}

	ctx.JSON(goods)
}

func GetGoodsDetail(ctx iris.Context){
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))

	goods := &models.Goods{}

	models.Db.Where("id = ?", goods_id).Find(goods)

	models.Db.Where("id = ?", goods.UserID).Find(&goods.User)

	models.Db.Where("goods_id = ?", goods_id).Find(&goods.Images)

	ctx.JSON(goods)
}
