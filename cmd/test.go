package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

var (
	openRtpServer = "/index/api/openRtpServer" //创建GB28181 RTP接收端口，如果该端口接收数据超时，则会自动被回收(不用调用closeRtpServer接口)
	getMediaList  = "/index/api/getMediaList"  //获取流媒体列表
)

type Params struct {
	Secret   string `json:"secret"`
	Port     int    `json:"port"`
	TcpMode  int    `json:"tcp_mode"`
	StreamId string `json:"stream_id"`
}

// 将结构体转换为 url.Values
func (p Params) toURLValues() url.Values {
	vals := url.Values{}
	val := reflect.ValueOf(p)
	typeOf := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typeOf.Field(i).Tag.Get("json")
		switch field.Kind() {
		case reflect.String:
			vals.Add(fieldName, field.String())
		case reflect.Int:
			vals.Add(fieldName, fmt.Sprintf("%d", field.Int()))
		}
	}
	return vals
}

func main() {
	// params := Params{
	// 	Secret:   "kHxDRQsk2xGwz1kvtkiH4BBt0rCs9oc9",
	// 	Port:     0,
	// 	TcpMode:  1,
	// 	StreamId: "0000059682",
	// }

	params := url.Values{}
	params.Add("secret", "kHxDRQsk2xGwz1kvtkiH4BBt0rCs9oc9")
	params.Add("port", "0")
	params.Add("tcp_mode", "1")
	params.Add("stream_id", "0000059682")

	url := fmt.Sprintf("http://127.0.0.1:28080%s?%s", openRtpServer, params.Encode())
	// url := fmt.Sprintf("http://127.0.0.1:28080%s?%s", getMediaList, params.Encode())

	fmt.Println(url)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making HTTP request:", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println(string(body))
}
