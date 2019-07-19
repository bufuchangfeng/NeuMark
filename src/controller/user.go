package controller

import (
	"../lib"
	"fmt"
	"github.com/kataras/iris"
)

func Bind(ctx iris.Context){
	sid := ctx.FormValue("sid")
	pwd := ctx.FormValue("pwd")

	fmt.Println(sid, pwd)

	lib.Login(sid, pwd)
}