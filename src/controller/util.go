package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/kataras/iris"
	"sort"
)

func CheckSignature(ctx iris.Context){
	signature := ctx.FormValue("signature")
	timestamp := ctx.FormValue("timestamp")
	nonce := ctx.FormValue("nonce")
	echostr := ctx.FormValue("echostr")

	token := "hellomarktoken"

	arr := []string{token, timestamp, nonce}
	sort.Strings(arr)

	tmpstr := arr[0] + arr[1] + arr[2]
	tmpstr = Sha1(tmpstr)

	if tmpstr == signature {
		ctx.WriteString(echostr)
	}
}


func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}
