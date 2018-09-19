package requests

import (
	"fmt"
	"testing"
)

func TestContent(t *testing.T) {
	url := "https://github.com"
	resp, err := Get(url, &Args{})
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}

	Content, err := resp.Content()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(Content[0])
	//fmt.Println(resp.content)
}

func TestText(t *testing.T) {
	url := "https://httpbin.org"
	resp, err := Get(url, &Args{})
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}

	resp.Text()
}

func TestJson(t *testing.T) {
	url := "https://httpbin.org/json"
	var data map[string]interface{}

	resp, err := Get(url, &Args{})
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}

	if resp.isOk(resp.StatusCode) {
		fmt.Println("Request successed")
	}
	resp.Json(&data)

	for k, v := range data {
		fmt.Println(k, v)
	}
}
