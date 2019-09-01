package models

import (
	"fmt"
	"github.com/jinzhu/gorm"

	_"github.com/jinzhu/gorm/dialects/mysql"
	// _"github.com/jinzhu/gorm/dialects/sqlite"

)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:yuchen@tcp(39.106.85.229:3306)/neumark?&parseTime=true&loc=Local")
	if nil != err {
		fmt.Println(err.Error())
	}
	return
}


// 数据库的外键可能有问题，记录一下。
//