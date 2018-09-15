[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/asmcos/requests/master/LICENSE)

# requests
Go HTTP Requests for Humans

## Installation
```
go get -u github.com/onesafe/requests
```

## Usage
```go
import (
    "github.com/onesafe/requests"
)
```

### GET
```go
resp, err := requests.Get("https://github.com")

if !resp.isOk(resp.StatusCode) {
	return errors.New("Get Failed")
}
// content type: byte[]
content, err := resp.Content()
// text type: string
text, err := resp.Text()
// data type can be: map[string]interface{}
err = resp.Json(&data)
```

### Set Headers
```go
url := "https://www.baidu.com"
	
h := Headers{
	"Referer":           "http://github.com",
	"Accept-Language":   "zh-CN,zh;",
	"Content-Type":      requests.ContentTypeJsonType,
}
resp, err := requests.Get(url, h)
```