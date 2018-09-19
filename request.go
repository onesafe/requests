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
	*Args
}

// Args contains all request arg
type Args struct {
	Headers HEADERS
	Params  PARAMS
	Datas   DATAS
}

type HEADERS map[string]string
type PARAMS map[string]string
type DATAS map[string]string

var (
	client *http.Client
	req    *http.Request
)

func NewRequest() *Request {
	if client == nil {
		client = new(http.Client)
	}

	args := &Args{}
	return &Request{Client: client, Args: args}
}

// GET
func Get(url string, args *Args) (resp *Response, err error) {
	resp, err = request("GET", url, args)
	return
}

func (r *Request) Get(url string) (resp *Response, err error) {
	if resp, err = r.coreRequest("GET", url); err != nil {
		fmt.Println("core Request failed")
	}
	return
}

// POST
func Post(url string, args *Args) (resp *Response, err error) {
	resp, err = request("POST", url, args)
	return
}

func (r *Request) Post(url string) (resp *Response, err error) {
	if resp, err = r.coreRequest("POST", url); err != nil {
		fmt.Println("core Request failed")
	}
	return
}

func request(method string, url string, args *Args) (resp *Response, err error) {
	r := NewRequest()
	if r.Client == nil {
		fmt.Println("init request client failed")
	}

	if resp, err = r.coreRequest(method, url); err != nil {
		fmt.Println("core Request failed")
	}
	return
}

func (r *Request) coreRequest(method string, url string) (resp *Response, err error) {
	if err = r.buildHTTPRequest(method, url); err != nil {
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

func (r *Request) buildHTTPRequest(method string, url string) (err error) {
	var body io.Reader

	if body, err = r.buildBody(); err != nil {
		return err
	}

	if r.Req, err = http.NewRequest(method, url, body); err != nil {
		return err
	}

	buildHeaders(r)
	buildURLParams(r, url)

	return
}

func (r *Request) buildBody() (body io.Reader, err error) {
	datas := map[string]string{}

	if r.Datas == nil {
		return nil, nil
	}
	datas = r.Datas

	// build post Form data
	Forms := url.Values{}

	for key, value := range datas {
		Forms.Add(key, value)
	}

	// build body
	body = strings.NewReader(Forms.Encode())
	return body, err
}

func (r *Request) SetTimeout(n time.Duration) {
	r.Client.Timeout = time.Duration(n * time.Second)
}
