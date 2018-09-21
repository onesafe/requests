package requests

import (
	"fmt"
	"testing"
)

func TestCookies(t *testing.T) {
	r := NewRequest()
	r.Cookies = COOKIES{
		"key": "value",
		"a": "123",
	}

	resp, _ := r.Get("https://httpbin.org/cookies")
	fmt.Println(resp.isOk(resp.StatusCode))
	fmt.Println(resp.toString())
}