package controller

import (
	"github.com/kataras/iris"
	"../models"
	"strconv"

)

func AddComment(ctx iris.Context){
	avatar := ctx.FormValue("avatar")
	nickname := ctx.FormValue("nickname")
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))
	content := ctx.FormValue("content")
	xid, _ := strconv.Atoi(ctx.FormValue("xid"))

	comment := &models.Comment{Avatar:avatar, NickName:nickname, GoodsID:goods_id, Content:content, XID:xid}

	models.Db.Create(comment)
}

func GetComments(ctx iris.Context) {
	goods_id, _ := strconv.Atoi(ctx.FormValue("goods_id"))
	var comments []models.Comment
	models.Db.Where("goods_id = ?", goods_id).Find(&comments)

	for i := 0; i < len(comments); i++ {
		if comments[i].XID != -1 {
			c := &models.Comment{ID:comments[i].XID}
			models.Db.First(c)

			comments[i].Xuser = c.NickName
			comments[i].Xcontent = c.Content
		}
	}

	ctx.JSON(comments)
}