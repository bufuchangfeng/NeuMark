package models

//商品类别
//类别id
//父类别id
//类别


// 这里有问题
type Category struct{
	ID int `gorm:"primary_key"`
	Name string
	CategoryID int
}
