package neu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"strings"
	"../../models"
)

func NEULogin(id, password string)(u models.User) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", "https://webvpn.neu.edu.cn", nil)
	if nil != err {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if nil != err {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var reg1 = regexp.MustCompile("input type=\"hidden\" id=\"lt\" name=\"lt\" value=\"(.*?)\" />")
	var reg2 = regexp.MustCompile("input type=\"hidden\" name=\"execution\" value=\"(.*?)\" />")

	lt := reg1.FindAllStringSubmatch(string(body), -1)[0][1]
	execution := reg2.FindAllStringSubmatch(string(body), -1)[0][1]

	str := "rsa=" + id + password + lt + "&ul=" + strconv.Itoa(len(id)) + "&pl=" + strconv.Itoa(len(password)) + "&lt=" + lt + "&execution=" + execution + "&_eventId=submit"
	data := strings.NewReader(str)

	req, err = http.NewRequest("POST", "https://pass.neu.edu.cn/tpass/login?service=https%3A%2F%2Fwebvpn.neu.edu.cn%2Fusers%2Fauth%2Fcas%2Fcallback%3Furl", data)
	if nil != err {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	res, err = client.Do(req)
	if nil != err {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	req, err = http.NewRequest("GET", "https://pass-443.webvpn.neu.edu.cn/tpass/login?service=http%3A%2F%2F219-216-96-4.webvpn.neu.edu.cn%2Feams%2FhomeExt.action", nil)
	if nil != err {
		fmt.Println(err)
		return
	}
	res, err = client.Do(req)
	if nil != err {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	lt = reg1.FindAllStringSubmatch(string(body), -1)[0][1]
	execution = reg2.FindAllStringSubmatch(string(body), -1)[0][1]

	str = "rsa=" + id + password + lt + "&ul=" + strconv.Itoa(len(id)) + "&pl=" + strconv.Itoa(len(password)) + "&lt=" + lt + "&execution=" + execution + "&_eventId=submit"
	data = strings.NewReader(str)

	req, err = http.NewRequest("POST", "https://pass-443.webvpn.neu.edu.cn/tpass/login?service=http%3A%2F%2F219-216-96-4.webvpn.neu.edu.cn%2Feams%2FhomeExt.action", data)
	if nil != err {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	res, err = client.Do(req)
	if nil != err {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()


	req, err = http.NewRequest("GET", "https://219-216-96-4.webvpn.neu.edu.cn/eams/stdDetail.action?", nil)
	if nil != err {
		fmt.Println(err)
		return
	}
	res, err = client.Do(req)
	if nil != err {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, _ = ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	reg3 := regexp.MustCompile("<td>(.*?)</td>")
	r := reg3.FindAllStringSubmatch(string(body), -1)

	// 错误处理函数 比如登录失败

	if len(r) >= 18  {
		name := r[0][1]
		sex := r[2][1]
		grade := r[3][1]
		institute := r[8][1]
		major := r[9][1]
		campus := r[17][1]
		class := r[18][1]

		tempuser := models.User{}

		models.Db.Where("sid = ?", id).Find(&tempuser)

		if tempuser.Name != "" && tempuser.Grade != ""{
			return tempuser
		} else {
			user := models.User{Name:name, Sex:sex, Sid:id, Grade:grade, Institute:institute, Major:major, Campus:campus, Class:class, Usertype:"1"}
			models.Db.Create(&user)
			fmt.Println(user)
			return user
		}
	}
	return
}
