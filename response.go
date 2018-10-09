package requests

import (
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
	content []byte
}

// Content return Response Body as []byte
func (resp *Response) Content() (b []byte, err error) {
	if resp.content != nil {
		return resp.content, nil
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		if reader, err = gzip.NewReader(resp.Body); err != nil {
			return nil, err
		}
	case "deflate":
		if reader, err = zlib.NewReader(resp.Body); err != nil {
			return nil, err
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()

	if resp.content, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return resp.content, nil
}

// Text return Response Body as string
func (resp *Response) Text() (text string, err error) {
	if resp.content == nil {
		_, err = resp.Content()
	}
	text = string(resp.content)
	return
}

// Json return Response Body as Json
func (resp *Response) Json(v interface{}) (err error) {
	if resp.content == nil {
		_, err = resp.Content()
	}
	return json.Unmarshal(resp.content, v)
}

// Check request success or fail
func (resp *Response) IsOk(code int) bool {
	switch code {
	case 200, 201, 202:
		return true
	}
	return false
}

// Same as Text() func, but toString func ignore error, it is more easy to test
func (resp *Response) ToString() (text string) {
	text = Fn(resp.Text)
	return
}

func Fn(f func() (string, error)) (res string) {
	res, _ = f()
	return
}
