package requests

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"

	args := &Args{
		TimeOut: 10,
	}
	resp, err := Get(url, args)
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	if !resp.isOk(resp.StatusCode) {
		t.Error("Simple Get Failed")
	}
}

func TestPost(t *testing.T) {
	url := "https://www.httpbin.org/post"

	data := DATAS{
		"name": "post_test",
	}

	args := &Args{
		Datas: data,
	}

	resp, _ := Post(url, args)
	d, _ := resp.Text()
	fmt.Println(d)
}
