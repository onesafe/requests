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
r := requests.NewRequest()

resp, err := r.Get("https://github.com")
if !resp.isOk(resp.StatusCode) {
	return errors.New("Get Failed")
}
```

### POST
```go
r := requests.NewRequest()

r.Datas = requests.DATAS{
	"name": "test"
}
resp, err := r.Post("https://www.httpbin.org/post")
```

### Response
```go
// content type: byte[]
content, err := resp.Content()

// text type: string
text, err := resp.Text()

// data type can be: map[string]interface{}
var data map[string]string
err = resp.Json(&data)
```

# Feature Support
  - Set Headers
  - Set Params
  - Set TimeOut
  - Auth


### Set Headers
```go
r := requests.NewRequest()

r.Headers = requests.HEADERS{
	"Referer":           "http://github.com",
	"Accept-Language":   "zh-CN,zh;",
	"Content-Type":      requests.ContentTypeJsonType,
}
resp, err := r.Get("https://www.baidu.com")
```

### Set Params
```go
r := requests.NewRequest()

r.Params = requests.PARAMS{
	"user":	"onesafe",
}
resp, err := r.Get("https://github.com")
```

### Set TimeOut
```go
r := requests.NewRequest()

r.SetTimeout(10) // 10 Seconds
resp, err := r.Get("https://github.com")
```

### Auth
```go
r := requests.NewRequest()

r.BasicAuth = requests.BasicAuth{"user", "passwd"}
resp, err := r.Get("http://httpbin.org/basic-auth/user/passwd")
```

