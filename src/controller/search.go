package controller

import (
	"github.com/kataras/iris"
	"../models"
	"strconv"
)
func Search(ctx iris.Context){
	content := ctx.FormValue("content")
	page := ctx.FormValue("page")
	page_size := ctx.FormValue("page_size")

	pageint, _ := strconv.Atoi(page)
	page_sizeint, _ := strconv.Atoi(page_size)

	content = "%" + content
	content = content +"%"

	var goods []models.Goods

	models.Db.Where("name like  ? OR description like ?", content).Offset(pageint * page_sizeint).Limit(page_size).Order("id desc").Find(&goods)

	for i := 0; i < len(goods); i++ {
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Comments)
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Collects)
	}

	ctx.JSON(goods)
}
