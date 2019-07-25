package hdu

import(
"fmt"
"github.com/robertkrimen/otto"
"io/ioutil"
"net/http"
"net/http/cookiejar"
"regexp"
"strconv"
"strings"
)

func HDULogin() {
	id := "18011628"
	password := "Hdu1117013"

	filePath := "des.js"

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err!=nil {
		panic(err)
	}

	//value, err := vm.Call("strEnc", nil, "abc", "1", "2", "3")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(value.String())

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", "https://cas.hdu.edu.cn/cas/login", nil)
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

	value, err := vm.Call("strEnc", nil, id + password + lt, "1", "2", "3")
	if err != nil {
		panic(err)
	}

	rsa, _ := value.ToString()

	str := "rsa=" + rsa + "&ul=" + strconv.Itoa(len(id)) + "&pl=" + strconv.Itoa(len(password)) + "&lt=" + lt + "&execution=" + execution + "&_eventId=submit"
	data := strings.NewReader(str)

	req, err = http.NewRequest("POST", "https://cas.hdu.edu.cn/cas/login?service=https%3A%2F%2Fi.hdu.edu.cn%2Ftp_up%2F", data)
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

	body, _ = ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))

	var reg3 = regexp.MustCompile("src=\"https://i.hdu.edu.cn/tp_up/\"><span class=\"tit\">(.*?)</span></a>")

	name := reg3.FindAllStringSubmatch(string(body), -1)[0][1]

	fmt.Println(name)
}