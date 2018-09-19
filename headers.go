package requests

var Version string = "1.0"
var DefaultUserAgent = "go-request/" + Version

var DefaultHeaders = map[string]string{
	"Connection":      "keep-alive",
	"Accept-Encoding": "gzip, deflate",
	"Accept":          "*/*",
	"User-Agent":      DefaultUserAgent,
}

var ContentTypeFormURLEncodedType = "application/x-www-form-urlencoded; charset=utf-8"
var ContentTypeJsonType = "application/json; charset=utf-8"

// construct Headers for request
func buildHeaders(r *Request) {

	//Set default value to Headers
	for k, v := range DefaultHeaders {
		r.Req.Header.Set(k, v)
	}

	if r.Headers == nil {
		return
	}

	for k, v := range r.Headers {
		r.Req.Header.Set(k, v)
	}
}
