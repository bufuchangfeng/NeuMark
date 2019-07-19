package controller

import (
	"github.com/kataras/iris"
	"../models"
)

func GetParentCategory(ctx iris.Context){
	var categories []models.Category

	models.Db.Find(&categories, "parent_id = ?", -1)
	ctx.JSON(categories)
}

func GetSubCategory(ctx iris.Context){
	parentId:= ctx.FormValue("parent_id")

	var categories []models.Category

	models.Db.Find(&categories, "parent_id = ?", parentId)
	ctx.JSON(categories)
}
