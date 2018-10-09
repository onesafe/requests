package requests

import (
	"net/http"
	"net/http/cookiejar"
)

func buildCookies(r *Request) {
	if r.Cookies == nil {
		return
	}

	// Check Client.Jar existed or not
	if r.Client.Jar == nil {
		jar, _ := cookiejar.New(nil)
		r.Client.Jar = jar
	}
	cookies := r.Client.Jar.Cookies(r.Req.URL)

	for key, value := range r.Cookies {
		cookies = append(cookies, &http.Cookie{Name: key, Value: value})
	}

	r.Client.Jar.SetCookies(r.Req.URL, cookies)
}