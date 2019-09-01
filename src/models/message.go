package models


type Message struct {
	ID int `gorm:primary_key`
	Avatar string
	NickName string
	GoodsID int
	Content string
	XUserID int
}
