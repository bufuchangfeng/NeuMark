package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("sqlite3", "neumark.db")
	if nil != err {
		fmt.Println(err.Error())
	}
	return
}


// 数据库的外键可能有问题，记录一下。
//