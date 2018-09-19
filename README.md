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
// if no args need to specify, you can directly use requests.Get
resp, _ := requests.Get("https://github.com")
text, _ := resp.Text()
```

```go
r := NewRequest()
resp, err := r.Get("https://github.com")

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
r := NewRequest()

r.Headers := requests.HEADERS{
	"Referer":           "http://github.com",
	"Accept-Language":   "zh-CN,zh;",
	"Content-Type":      requests.ContentTypeJsonType,
}
resp, err := r.Get("https://www.baidu.com")
```

### Set Params
```go
r := NewRequest()

r.Params := requests.PARAMS{
	"user":	"onesafe",
}
resp, err := r.Get("https://github.com")
```
