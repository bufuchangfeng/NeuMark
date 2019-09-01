package controller

import (
	"../models"
	"github.com/kataras/iris"
	"strconv"
)

func GetMessages(ctx iris.Context){
	userid, _ := strconv.Atoi(ctx.FormValue("userid"))
	var messgaes []models.Message
	models.Db.Where("x_user_id = ?", userid).Find(&messgaes)

	ctx.JSON(messgaes)
}

func DelMessage(ctx iris.Context){
	id, _ := strconv.Atoi(ctx.FormValue("id"))
	models.Db.Where("id", id).Delete(&models.Message{})
}
