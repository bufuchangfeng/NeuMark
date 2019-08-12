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
