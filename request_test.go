package requests

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"

	var n TimeOut = 10
	resp, err := Get(url, n)
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
