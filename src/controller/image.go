package controller

import (
	"github.com/kataras/iris"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"../models"
	"strconv"
)

func AddImage(ctx iris.Context) {
	GoodsID := ctx.FormValue("goods_id")
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	form := ctx.Request().MultipartForm

	files := form.File["files[]"]
	failures := 0
	for _, file := range files {
		_, err = saveUploadedFile(file, "../goods", GoodsID)
		if err != nil {
			failures++
			ctx.Writef("failed to upload: %s\n", file.Filename)
		}
	}
	ctx.Writef("%d files uploaded", len(files)-failures)
}


func saveUploadedFile(fh *multipart.FileHeader, destDirectory string, goodsID string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, goodsID + "-" + fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))

	if err != nil {
		return 0, err
	}
	defer out.Close()

	id, _ := strconv.Atoi(goodsID)
	models.Db.Create(&models.Image{GoodsID:id, Filename:goodsID + "-" + fh.Filename})

	return io.Copy(out, src)
}

