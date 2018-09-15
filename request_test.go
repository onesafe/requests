package requests

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"
	resp, err := Get(url)
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	if resp.StatusCode != 200 {
		t.Error("Simple Get Failed")
	}
}