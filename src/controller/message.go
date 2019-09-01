package controller

import (
	"../models"
	"github.com/kataras/iris"
	"strconv"
	"fmt"
)

func GetMessages(ctx iris.Context){
	userid, _ := strconv.Atoi(ctx.FormValue("userid"))
	fmt.Println(userid)
	var messages []models.Message
	models.Db.Where("x_user_id = ?", userid).Order("id desc").Find(&messages)

	fmt.Println(messages)

	ctx.JSON(messages)
}

func DelMessage(ctx iris.Context){
	id, _ := strconv.Atoi(ctx.FormValue("id"))
	models.Db.Where("id = ?", id).Delete(&models.Message{})
}
