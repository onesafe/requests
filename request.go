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
	Headers   HEADERS
	Params    PARAMS
	Datas     DATAS
	BasicAuth BasicAuth
	Proxy     string
}

type HEADERS map[string]string
type PARAMS map[string]string
type DATAS map[string]string

// BasicAuth struct for http basic auth
type BasicAuth struct {
	Username string
	Password string
}

var (
	_CLIENT      *http.Client
	req          *http.Request
	maxIdleConns int = 10
)

func NewRequest() *Request {
	args := &Args{}

	if _CLIENT != nil {
		return &Request{Client: _CLIENT, Args: args}
	}

	transport := &http.Transport{
		MaxIdleConns:    maxIdleConns,
		IdleConnTimeout: 30 * time.Second,
	}
	_CLIENT = &http.Client{
		Transport: transport,
	}

	return &Request{Client: _CLIENT, Args: args}
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

// Below are Request Struct Functions
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

	if r.BasicAuth.Username != "" {
		r.SetBasicAuth(r.BasicAuth.Username, r.BasicAuth.Password)
	}

	buildHeaders(r)

	if err = buildURLParams(r, url); err != nil {
		fmt.Println("build URL Params Failed" + err.Error())
	}

	if err = buildProxy(r); err != nil {
		fmt.Println("Set Proxy Failed" + err.Error())
	}

	return
}

func (r *Request) buildBody() (body io.Reader, err error) {
	if r.Datas == nil {
		return nil, nil
	}

	// build post Form data
	Forms := url.Values{}
	for key, value := range r.Datas {
		Forms.Add(key, value)
	}

	// build body
	body = strings.NewReader(Forms.Encode())
	return body, err
}

// Set Request TimeOut
func (r *Request) SetTimeout(n time.Duration) {
	r.Client.Timeout = time.Duration(n * time.Second)
}

// Set Basic Auth
func (r *Request) SetBasicAuth(Username string, Password string) {
	r.Req.SetBasicAuth(Username, Password)
}

// Reset client and Args to default values
func (r *Request) Reset() {
	_CLIENT = nil

	r.Headers = nil
	r.Params = nil
	r.Datas = nil
	r.Proxy = ""
	r.BasicAuth = BasicAuth{}
	return
}
