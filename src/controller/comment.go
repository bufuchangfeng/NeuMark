package controller

import (
	"github.com/kataras/iris"
	"../models"
	"strconv"
	"time"
)

func AddComment(ctx iris.Context){
	avatar := ctx.FormValue("avatar")
	nickname := ctx.FormValue("nickname")
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))
	content := ctx.FormValue("content")
	xid, _ := strconv.Atoi(ctx.FormValue("xid"))

	comment := &models.Comment{Avatar:avatar, NickName:nickname, GoodsID:goods_id, Content:content, XID:xid, CreatedAt:time.Now()}

	models.Db.Create(comment)
}
