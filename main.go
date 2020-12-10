package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)
var (str string)

func main() {
    str = fmt.Sprintf("text=测试信息,Action来自 【 %s 】",os.Args[1])
	fmt.Println(str)
	//	send()

}

func send() {
	url := "http://sc.ftqq.com/SCU124621T9af583eb31c71ee1ac579fe2db1fbeab5fa7fcd8a0079.send"
	method := "POST"

	payload := strings.NewReader(str)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Connection", "keep-alive")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

