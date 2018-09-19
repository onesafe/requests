package requests

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"

	args := &Args{}
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

func TestRequestGet(t *testing.T) {
	url := "https://www.baidu.com"

	r := NewRequest()

	resp, _ := r.Get(url)
	if !resp.isOk(resp.StatusCode) {
		t.Error("Simple Get Failed")
	}

	r.Reset()
	fmt.Println(resp.Status)
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

func TestRequestPost(t *testing.T) {
	url := "https://www.httpbin.org/post"

	r := NewRequest()
	r.Datas = DATAS{
		"name": "post_test",
	}

	r.SetTimeout(20)

	resp, err := r.Post(url)
	if err != nil {
		fmt.Println("Post error: " + err.Error())
	}
	d, _ := resp.Text()
	fmt.Println(d)
}

func TestBasicAuth(t *testing.T) {
	r := NewRequest()
	r.BasicAuth = BasicAuth{"user", "passwd"}
	resp, err := r.Get("http://httpbin.org/basic-auth/user/passwd")
	if err != nil {
		t.Error(err.Error())
	}
	d, _ := resp.Text()
	fmt.Println(d)
}
