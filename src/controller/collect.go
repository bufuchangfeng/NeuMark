package controller

import (
	"github.com/kataras/iris"
	"strconv"
	"../models"
	"time"
)

func AddCollect(ctx iris.Context){
	user_id, _ := strconv.Atoi(ctx.FormValue("user_id"))
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))

	collect := &models.Collect{UserID:user_id, GoodsID:goods_id}
	if !(collect.ID > 0){
		collect = &models.Collect{UserID:user_id, GoodsID:goods_id, CreatedAt:time.Now()}
		models.Db.Create(collect)
	}
}

func GetCollects(ctx iris.Context){
	user_id, _ := strconv.Atoi(ctx.FormValue("user_id"))

	var collects []models.Collect

	models.Db.Where("user_id = ?", user_id).Find(&collects)

	var goods [20]models.Goods
	for i := 0; i < len(collects); i++ {
		models.Db.Where("goods_id = ?", collects[i].GoodsID).Find(&goods[i])
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Comments)
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Collects)
	}

	ctx.JSON(goods)
}