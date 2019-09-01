package models

import "time"

//评论id
//商品id
//发布人id
//评论内容

type Comment struct{
	ID int `gorm:primary_key`
	Avatar string
	NickName string
	GoodsID int
	Content string
	Xuser string
	Xcontent string
	XID	int
	UserID int
	CreatedAt time.Time
}
