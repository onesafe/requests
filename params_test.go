package requests

import (
	"fmt"
	"testing"
)

func TestParams(t *testing.T) {
	p := Params{
		"user": "onesafe",
	}

	resp, err := Get("https://github.com", p)
	if err != nil {
		t.Error("[GET] failed " + err.Error())
	}

	_, err = resp.Content()
	fmt.Println("Test Params Start")
	fmt.Println(resp.Status)
	// fmt.Println(data)
	fmt.Println("Test Params end")
}