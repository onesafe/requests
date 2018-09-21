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
	Hooks     []Hook
	Cookies   COOKIES
}

type HEADERS map[string]string
type PARAMS map[string]string
type DATAS map[string]string
type COOKIES map[string]string

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
		fmt.Println("[GET] core Request failed")
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
		fmt.Println("[POST] core Request failed")
	}
	return
}

// PUT
func Put(url string, args *Args) (resp *Response, err error) {
	resp, err = request("PUT", url, args)
	return
}

func (r *Request) Put(url string) (resp *Response, err error) {
	if resp, err = r.coreRequest("PUT", url); err != nil {
		fmt.Println("[PUT] core Request failed")
	}
	return
}

// PATCH
func Patch(url string, args *Args) (resp *Response, err error) {
	resp, err = request("PATCH", url, args)
	return
}

func (r *Request) Patch(url string) (resp *Response, err error) {
	if resp, err = r.coreRequest("PATCH", url); err != nil {
		fmt.Println("[PATCH] core Request failed")
	}
	return
}

// DELETE
func Delete(url string, args *Args) (resp *Response, err error) {
	resp, err = request("DELETE", url, args)
	return
}

func (r *Request) Delete(url string) (resp *Response, err error) {
	if resp, err = r.coreRequest("DELETE", url); err != nil {
		fmt.Println("[DELETE] core Request failed")
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

	// apply BeforeRequest Hook
	s, err := applyBeforeReqHooks(r.Req, r.Hooks)
	if err != nil {
		return nil, err
	} else if s != nil {
		resp = &Response{s, nil}
		return resp, err
	}

	res, err := r.Client.Do(r.Req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// apply AfterRequest hook
	newResp, newErr := applyAfterReqHooks(r.Req, res, err, r.Hooks)
	if newErr != nil {
		err = newErr
	}
	if newResp != nil {
		res = newResp
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

	buildCookies(r)

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

// Set Connection Pool Size
func (r *Request) SetPoolSize(size int) {
	r.Client.Transport = &http.Transport{
		MaxIdleConns:    size,
		IdleConnTimeout: 30 * time.Second,
	}
}

// Reset client and Args to default values
func (r *Request) Reset() {
	_CLIENT = nil

	r.Headers = nil
	r.Params = nil
	r.Datas = nil
	r.Proxy = ""
	r.BasicAuth = BasicAuth{}
	r.Cookies = nil
	r.Hooks = []Hook{}
	return
}
