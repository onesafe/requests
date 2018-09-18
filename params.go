package requests

import (
	"net/url"
	"strings"
)

func buildURLParams(r *Request, URL string, args ...interface{}) (err error) {
	params := []map[string]string{}
	var tempUrl string

	// Get params from args
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

	// get original parse query ?user=onesafe
	// parsedQuery type: map[user:[onesafe]]
	parsedQuery, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return err
	}

	// add params to Query
	for _, param := range params {
		for key, value := range param {
			parsedQuery.Add(key, value)
		}
	}

	// remove original query param
	withoutQueryUrl := strings.Replace(parsedURL.String(), "?"+parsedURL.RawQuery, "", -1)

	if len(parsedQuery) > 0 {
		// add original query param and args params
		tempUrl = strings.Join([]string{withoutQueryUrl, parsedQuery.Encode()}, "?")
	}
	tempUrl = withoutQueryUrl

	destUrl, err := url.Parse(tempUrl)
	if err != nil {
		return err
	}

	// URL with Query params
	r.Req.URL = destUrl
	return
}
