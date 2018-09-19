package requests

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	Req    *http.Request
	Client *http.Client
}

type Headers map[string]string
type Params map[string]string
type TimeOut time.Duration
type Datas map[string]string

var (
	client *http.Client
	req    *http.Request
)

func NewRequest() *Request {
	if client == nil {
		client = new(http.Client)
	}

	return &Request{Client: client}
}

// GET
func Get(url string, args ...interface{}) (resp *Response, err error) {
	resp, err = request("GET", url, args)
	return
}

// POST
func Post(url string, args ...interface{}) (resp *Response, err error) {
	resp, err = request("POST", url, args)
	return
}

func request(method string, url string, args ...interface{}) (resp *Response, err error) {
	r := NewRequest()
	if r.Client == nil {
		fmt.Println("init request client failed")
	}

	if err = r.buildHTTPRequest(method, url, args); err != nil {
		fmt.Println("build HTTP Request failed" + err.Error())
	}

	res, err := r.Client.Do(r.Req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp = &Response{res, nil}
	return
}

func (r *Request) buildHTTPRequest(method string, url string, args ...interface{}) (err error) {
	var body io.Reader

	if body, err = r.buildBody(args); err != nil {
		return err
	}

	if r.Req, err = http.NewRequest(method, url, body); err != nil {
		return err
	}

	buildHeaders(r, args)
	buildURLParams(r, url, args)

	SetTimeout(r, args)

	return
}

func (r *Request) buildBody(args ...interface{}) (body io.Reader, err error) {
	datas := []map[string]string{}

	for _, arg := range args {
		switch customType := arg.(type) {
		case Datas:
			datas = append(datas, customType)
		}
	}

	// build post Form data
	Forms := url.Values{}
	for _, data := range datas {
		for key, value := range data {
			Forms.Add(key, value)
		}
	}

	// build body
	body = strings.NewReader(Forms.Encode())
	return body, err
}

func SetTimeout(r *Request, args ...interface{}) {
	var n time.Duration
	for _, arg := range args {
		switch customType := arg.(type) {
		case TimeOut:
			n = time.Duration(customType)
		}
	}
	r.Client.Timeout = time.Duration(n * time.Second)
}
