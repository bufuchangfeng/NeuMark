package neu

import (
	"../../models"
)

func ShopBind(id, password string)(u models.User){
	tempuser := models.User{}

	models.Db.Where("name = ? and password = ?", id, password).Find(&tempuser)

	return tempuser
}
