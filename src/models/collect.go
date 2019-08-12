package models

import "time"

// 后期加入
type Collect struct{
	ID int `gorm:primary_key`
	UserID int
	GoodsID int
	CreatedAt time.Time
}
