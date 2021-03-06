package controller

import (
	"../models"
	"fmt"
	"github.com/kataras/iris"
	"strconv"
)

func DeleteCollect(ctx iris.Context){
	user_id, _ := strconv.Atoi(ctx.FormValue("user_id"))
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))


	models.Db.Where("goods_id = ? AND user_id = ?", goods_id, user_id).Delete(&models.Collect{})
}


func AddCollect(ctx iris.Context){
	user_id, _ := strconv.Atoi(ctx.FormValue("user_id"))
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))

	collect := &models.Collect{}
	models.Db.Where("user_id = ? AND goods_id = ?", user_id, goods_id).Find(collect)
	fmt.Println(collect)
	if !(collect.ID > 0){
		collect = &models.Collect{UserID:user_id, GoodsID:goods_id}
		models.Db.Create(collect)
	}
}

func GetCollects(ctx iris.Context){
	user_id, _ := strconv.Atoi(ctx.FormValue("user_id"))

	var collects []models.Collect

	models.Db.Where("user_id = ?", user_id).Find(&collects)

	fmt.Println(collects)

	var goods [30]models.Goods
	for i := 0; i < len(collects); i++ {
		models.Db.Where("id = ?", collects[i].GoodsID).Find(&goods[i])

		fmt.Println(goods[i])
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Comments)
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Collects)
	}

	ctx.JSON(goods)
}