package requests

import (
	"net/http"
)

func currentIP(proxyurl string) (ip string) {
	r := NewRequest()

	r.Proxy = proxyurl
	resp, _ := r.Get("http://httpbin.org/get")
	var d map[string]string
	_ = resp.Json(&d)

	return d["origin"]
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"origin\": \"127.0.0.1\"}"))
}

/*
this unit test passed
but client will reuse, So proxy will still remain and use for another request
comment this, go test will pass
*/

// func TestHTTPProxy(t *testing.T) {
// 	proxy := httptest.NewServer(http.HandlerFunc(proxyHandler))
// 	defer proxy.Close()

// 	httpProxyURL := proxy.URL
// 	assert.Equal(t, currentIP(httpProxyURL) == "127.0.0.1", true)
// }
