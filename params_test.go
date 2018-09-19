package requests

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParams(t *testing.T) {
	p := PARAMS{
		"user": "onesafe",
	}
	args := &Args{
		Params: p,
	}

	resp, err := Get("https://github.com", args)
	if err != nil {
		t.Error("[GET] failed " + err.Error())
	}

	_, err = resp.Content()
	fmt.Println("Test Params Start")
	fmt.Println(resp.Status)
	// fmt.Println(data)
	fmt.Println("Test Params end")
}

func TestParseURL(t *testing.T) {
	URL := "https://github.com?user=onesafe"
	parsedURL, err := url.Parse(URL)
	if err != nil {
		t.Error("url parse error")
	}
	fmt.Println(parsedURL)

	fmt.Println(parsedURL.RawQuery)
	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		t.Error("url parse Query error")
	}
	fmt.Println(parsedQuery)
}
