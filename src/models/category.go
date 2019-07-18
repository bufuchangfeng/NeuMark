package models

//商品类别
//类别id
//父类别id
//类别

type Category struct{
	ID int `gorm:"primary_key"`
	Name string
	Category Category `gorm:"ForeignKey:CategoryId`
	CategoryId int
}
