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

### GET:
```go
resp, err := requests.Get("https://github.com")
if resp.StatusCode != 200 {
	return errors.New("Get Failed")
}
// content type: byte[]
content, err := resp.Content()
```
