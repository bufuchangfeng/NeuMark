package models

import "time"


// 后期加入
type Buy struct{
	ID int `gorm:primary_key`
	UserID int
	GoodsID int
	Time time.Time
}
