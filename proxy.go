package requests

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

func buildProxy(r *Request) (err error) {
	if r.Proxy == "" {
		return nil
	}

	u, err := url.Parse(r.Proxy)
	if err != nil {
		return err
	}

	switch u.Scheme {
	case "http", "https":
		r.Client.Transport = &http.Transport{
			Proxy: http.ProxyURL(u),
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
	}
	return
}
