package requests

import (
	"fmt"
	"testing"
)

func TestContent(t *testing.T) {
	url := "https://github.com"
	resp, err := Get(url)
	if err != nil {
		fmt.Println("[GET] Failed while request url: " + url)
	}

	Content, err := resp.Content()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(Content)
	fmt.Println(resp.content)
}