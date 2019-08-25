package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"../lib/neu"
	"../models"
	"strconv"
)

func Bind(ctx iris.Context){
	sid := ctx.FormValue("sid")
	pwd := ctx.FormValue("pwd")

	fmt.Println(sid, pwd)

	user := neu.NEULogin(sid, pwd)

	ctx.JSON(user)
}


func GetUserInfo(ctx iris.Context){
	user_id := ctx.FormValue("user_id")

	user := &models.User{}

	models.Db.Where("id = ?", user_id).Find(&user)

	ctx.JSON(user)
}

func UpdateUserInfo(ctx iris.Context){
	user_id := ctx.FormValue("user_id")
	wechat := ctx.FormValue("wechat")
	phone := ctx.FormValue("phone")
	qq := ctx.FormValue("qq")


	user_id_int, _ := strconv.Atoi(user_id)
	user := models.User{ID:user_id_int}
	models.Db.Where("id = ?", user_id_int).Find(&user)

	fmt.Println(user_id_int, user)

	models.Db.Model(&user).Updates(map[string]interface{}{"phone": phone, "we_chat": wechat, "qq": qq})
}

