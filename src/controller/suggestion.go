package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"../models"
	"strconv"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func AddSuggestion(ctx iris.Context){
	user_id, _ := strconv.Atoi(ctx.FormValue("user_id"))
	content := ctx.FormValue("content")

	models.Db.Create(&models.Suggestion{UserID:user_id, Content:content})

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "924761163@qq.com", "null")
	m.SetAddressHeader("To", "924761163@qq.com", "null")
	m.SetHeader("Subject", "Mark意见")
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.qq.com", 587, "924761163@qq.com", "sfpcogiarsnhbfgh")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
