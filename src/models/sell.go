package models

import "time"

type Sell struct {
	ID int `gorm:primary_key`
	UserID int
	GoodsID int
	CreatedAt time.Time
}