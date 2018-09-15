package requests

import (
	"compress/zlib"
	"compress/gzip"
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