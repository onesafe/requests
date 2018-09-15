package requests

import (
	"testing"
	"fmt"
)

func TestHeaders(t *testing.T) {
	url := "https://www.baidu.com"
	
	h := Headers{
		"Referer":           "http://github.com",
		"Accept-Language":   "zh-CN,zh;",
		"Content-Type":      ContentTypeJsonType,
	}
	resp, err := Get(url, h)
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}

	if !resp.isOk(resp.StatusCode) {
		t.Error("Simple Get Failed")
	}
}