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
	Category Category `gorm:"ForeignKey:CategoryId`
	CategoryId int
	Description string
	// images
	Price string
	View int
	Status string
	User User `gorm:"ForeignKey:UserId`
	UserId int
	CreatedAt time.Time
}
