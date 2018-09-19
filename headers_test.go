package requests

import (
	"fmt"
	"testing"
)

func TestHeaders(t *testing.T) {
	url := "https://www.baidu.com"

	h := HEADERS{
		"Referer":         "http://github.com",
		"Accept-Language": "zh-CN,zh;",
		"Content-Type":    ContentTypeJsonType,
	}

	args := &Args{
		Headers: h,
	}
	resp, err := Get(url, args)
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}

	if !resp.isOk(resp.StatusCode) {
		t.Error("Simple Get Failed")
	}
}
