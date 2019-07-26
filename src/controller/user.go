package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"../lib/neu"
)

func Bind(ctx iris.Context){
	sid := ctx.FormValue("sid")
	pwd := ctx.FormValue("pwd")

	fmt.Println(sid, pwd)

	user := neu.NEULogin(sid, pwd)

	ctx.JSON(user)
}