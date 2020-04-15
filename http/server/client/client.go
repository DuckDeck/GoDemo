package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	sendGetPara()
}

func sendGet() {
	res, err := http.Get("http://127.0.0.1:9999/get/peg?name=老王&age=123")
	if err != nil {
		fmt.Println("get 127.0.0.1:9999/get/peg fail err:", err)
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("red response fail err:", err)
		return
	}
	fmt.Println(string(b))
}

func sendGetPara() {
	urlObj, _ := url.Parse("http://127.0.0.1:9999/get/peg")
	data := url.Values{}
	data.Set("name", "小王")
	data.Set("age", "99993")
	urlStr := data.Encode()
	urlObj.RawQuery = urlStr

	req, err := http.NewRequest("GET", urlObj.String(), nil)

	if err != nil {
		fmt.Println("get 127.0.0.1:9999/get/peg fail err:", err)
		return
	}
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http send fail err:", err)
		return
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("red response fail err:", err)
		return
	}
	fmt.Println(string(b))
}
