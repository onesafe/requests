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

	if !resp.IsOk(resp.StatusCode) {
		t.Error("Simple Get Failed")
	}
}

func TestRequestGet(t *testing.T) {
	url := "https://www.baidu.com"

	r := NewRequest()
	r.SetPoolSize(50)

	resp, _ := r.Get(url)
	if !resp.IsOk(resp.StatusCode) {
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

func TestPut(t *testing.T) {
	r := NewRequest()
	url := "http://httpbin.org/put"

	r.Datas = DATAS{"put": "test"}
	resp, _ := r.Put(url)
	data, _ := resp.Text()
	fmt.Println(data)
}

func TestDelete(t *testing.T) {
	r := NewRequest()
	url := "http://httpbin.org/delete"

	resp, _ := r.Delete(url)
	data, _ := resp.Text()
	fmt.Println(data)
}

func TestPatch(t *testing.T) {
	r := NewRequest()
	url := "http://httpbin.org/patch"

	r.Datas = DATAS{"patch": "test"}
	resp, _ := r.Patch(url)
	data, _ := resp.Text()
	fmt.Println(data)
}
