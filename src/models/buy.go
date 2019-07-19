package models

type Buy struct{
	ID int `gorm:primary_key`
	UserID int
	Goods Goods `gorm:ForeignKey:GoodsID`
	GoodsID int
}
