package requests

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"net/http"
	"testing"
	"errors"
	"encoding/json"
)

type hookNothing struct {
	callBeforeHook bool
	callAfterHook bool
}

func (h *hookNothing) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	h.callBeforeHook = true
	return
}

func (h *hookNothing) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	h.callAfterHook = true
	return
}

func TestHookNothing(t *testing.T) {
	h := &hookNothing{}

	r := NewRequest()
	r.Hooks = []Hook{h}

	resp, _ := r.Get("https://httpbin.org/get")

	fmt.Println(resp.isOk(resp.StatusCode))
	fmt.Println(h.callBeforeHook)
	fmt.Println(h.callAfterHook)
}

type beforeRequestHookError struct {
	err error
}

func (h *beforeRequestHookError) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	err = h.err
	return
}

func (h *beforeRequestHookError) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	return
}

func TestBeforeRequestHookError(t *testing.T) {
	e := errors.New("before request hook error")
	h := &beforeRequestHookError{e}

	r := NewRequest()
	r.Hooks = []Hook{h}

	_, err := r.Get("https://httpbin.org/get")
	fmt.Println(err)
}

type beforeRequestHookResp struct {
	resp *http.Response
}

func (h *beforeRequestHookResp) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	resp = h.resp
	return
}

func (h *beforeRequestHookResp) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	return
}

func TestBeforeRequestHookResp(t *testing.T) {
	j, _ := json.Marshal(map[string]string{
		"url": "http://test",
	})

	b := ioutil.NopCloser(bytes.NewReader(j))

	s := &http.Response{Body: b}
	h := &beforeRequestHookResp{s}

	r := NewRequest()
	r.Hooks = []Hook{h}
	resp, _ := r.Get("https://httpbin.org/get")
	fmt.Println(resp.toString())
}