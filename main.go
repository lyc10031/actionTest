package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)
var (str,security  string)

func main() {
    str = fmt.Sprintf("text=测试信息,Action来自 【 %s 】",os.Args[1],)
	security = os.Args[2]
	fmt.Println(str,security)
		send()

}

func send() {
	url := "http://sc.ftqq.com/" + security + ".send"
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

