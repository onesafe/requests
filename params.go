package requests

import (
	"strings"
	"net/url"
)

func buildURLParams(r *Request, URL string, args ...interface{}) (err error) {
	params := []map[string]string{}
	var tempUrl string

	for _, arg := range args {
		switch customType := arg.(type) {
		case Params:
			params = append(params, customType)
		}
	}

	parsedURL, err := url.Parse(URL)
	if err != nil {
		return err
	}

	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return err
	}

	for _, param := range params {
		for key, value := range param {
			parsedQuery.Add(key, value)
		}
	}

	withoutQueryUrl := strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1)

	if len(parsedQuery) > 0 {
		tempUrl = strings.Join([]string{withoutQueryUrl, parsedQuery.Encode()}, "?")
	}
	tempUrl = withoutQueryUrl

	destUrl, err := url.Parse(tempUrl)
	if err != nil {
		return err
	}

	r.Req.URL = destUrl
	return
}