package controller

import (
	"github.com/kataras/iris"
	"../models"
	"strconv"
	"fmt"
	"io"
	"os"
)

func AddImage(ctx iris.Context) {
	GoodsID := ctx.FormValue("goods_id")
	index := ctx.FormValue("index")
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	file, info, _ := ctx.FormFile("files")
	fmt.Println(file, info)
	defer file.Close()

	out, err := os.OpenFile("../goods/"+ GoodsID + "-" + index + ".jpg", os.O_WRONLY|os.O_CREATE, 0666)

	defer out.Close()
	io.Copy(out, file)

	id, _ := strconv.Atoi(GoodsID)
	models.Db.Create(&models.Image{GoodsID:id, Filename:GoodsID + "-" + index + ".jpg"})
}
