package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	str, security string
	client        = &http.Client{}
	WxUrl         = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
)

func main() {
	security = os.Args[1]
	WxUrl += security
	send2(WxUrl)
}

type WxMsg struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}
type Text struct {
	Content string `json:"content"`
}

func send2(url string) {

	ms := WxMsg{}
	ms.Msgtype = "text"
	ms.Text = Text{}
	ms.Text.Content = "消息来自GitHubAction定时任务，\n【该写周报了~~】"

	data, err := json.Marshal(ms)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	resp, err := client.Post(WxUrl, "application/x-www-form-urlencoded", strings.NewReader(string(data)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	st, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("推送结果:\n\t ", string(st))
}
