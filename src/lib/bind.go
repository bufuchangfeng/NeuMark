package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"strings"
)

func Login(id, password string) {
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
	// fmt.Println(string(body))

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

	// body, _ = ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))
	// fmt.Println(res.StatusCode)

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

	// body, _ = ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

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

	// body, _ = ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

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
}
