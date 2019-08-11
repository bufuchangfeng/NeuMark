package models

import "time"

//商品id
//名称
//分类
//描述 文字描述 图片描述
//价格
//浏览量
//状态
//发布人学号
//发布时间

type Goods struct{
	ID int `gorm:primary_key`
	Name string
	Category Category `gorm:"ForeignKey:CategoryID`
	CategoryID int
	Description string
	images string
	Price string
	Hash string
	// View int
	// Status string
	User User
	UserID int
	CreatedAt time.Time

	Sells []Sell	`gorm:"FOREIGNKEY:GoodsID;ASSOCIATION_FOREIGNKEY:ID"`
	Comments []Comment `gorm:"FOREIGNKEY:GoodsID;ASSOCIATION_FOREIGNKEY:ID"`
	Buys []Buy		`gorm:"FOREIGNKEY:GoodsID;ASSOCIATION_FOREIGNKEY:ID"`
	Collects []Collect `gorm:"FOREIGNKEY:GoodsID;ASSOCIATION_FOREIGNKEY:ID"`

	Images []Image
}
