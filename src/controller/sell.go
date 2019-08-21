package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"strconv"
	"../models"
)

func GetSells(ctx iris.Context){
	user_id, _ := strconv.Atoi(ctx.FormValue("user_id"))

	var sells []models.Sell

	models.Db.Where("user_id = ?", user_id).Find(&sells)

	fmt.Println(sells)

	var goods [20]models.Goods
	for i := 0; i < len(sells); i++ {
		models.Db.Where("id = ?", sells[i].GoodsID).Find(&goods[i])

		fmt.Println(goods[i])
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Comments)
		models.Db.Where("goods_id = ?", goods[i].ID).Find(&goods[i].Collects)
	}

	ctx.JSON(goods)
}

func DeleteSell(ctx iris.Context){
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))

	models.Db.Where("goods_id = ?", goods_id).Delete(&models.Sell{})
	models.Db.Where("goods_id = ?", goods_id).Delete(&models.Collect{})
	models.Db.Where("id = ?", goods_id).Delete(&models.Goods{})
}
