package models

//评论id
//商品id
//发布人id
//评论内容

type Comment struct{
	ID int `gorm:primary_key`
	UserID int
	GoodsID int
	Description string
}
