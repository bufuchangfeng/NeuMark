package controller

import (
	"fmt"
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
	user_id, _ := strconv.Atoi(ctx.FormValue("userid"))


	l,_ := time.LoadLocation("Asia/Shanghai")
	comment := &models.Comment{Avatar:avatar, NickName:nickname, GoodsID:goods_id, Content:content, XID:xid, UserID:user_id,CreatedAt:time.Now().In(l)}
	models.Db.Create(comment)
	fmt.Println(time.Now().In(l))
	// 添加两条消息
	// 第一条是卖家
	// 根据GoodsID找到卖家的ID
	sell := &models.Sell{}
	models.Db.Where("goods_id = ?", goods_id).Find(sell)
	m := &models.Message{Avatar:avatar, NickName:nickname, GoodsID:goods_id, Content:content, XUserID:sell.UserID}
	models.Db.Create(m)

	// 第二条是被回复的人，如果这条评论是回复的话
	if xid != -1 {
		// 根据xid找到被回复用户的ID
		c := &models.Comment{ID:xid}
		models.Db.First(c)

		m := &models.Message{Avatar:avatar, NickName:nickname, GoodsID:goods_id, Content:content, XUserID:c.UserID}
		models.Db.Create(m)
	}
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