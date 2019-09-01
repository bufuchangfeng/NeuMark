package controller

import (
	"../models"
	"github.com/kataras/iris"
	"strconv"
)

func GetMessages(ctx iris.Context){
	// userid, _ := strconv.Atoi(ctx.FormValue("userid"))
	var messgaes []models.Message
	models.Db.Where("").Find(&messgaes)

}

func DelMessage(ctx iris.Context){

}
