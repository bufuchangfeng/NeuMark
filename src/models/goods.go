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
	// View int
	// Status string
	User User `gorm:"ForeignKey:UserID`
	UserID int
	CreatedAt time.Time

	Comments []Comment
}
