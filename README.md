# ogs-go

Easier to get a consistent structure

##### `RspOK(msgContent string)`

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/xandercheung/ogs-go"
)

func main() {
	v := ogs.RspOK("this is a message")
	b, _ := json.Marshal(v)
	fmt.Println(string(b))
}

/**
{
    "code": 0,
    "message": {
        "content": "this is a message",
        "type": "success"
    }
}
*/

```

##### `Rsp(code interface{}, message Message)`
```go
ogs.Rsp(1001, ogs.NewMsg("this is a message", "error"))

/**
{
    "code": 1001,
    "message": {
        "content": "this is a message",
        "type": "error"
    }
}
*/
```

#### `RspData(code interface{}, message Message, data interface{})`

```go
ogs.RspData(
    1002,
    ogs.WarningMsg("this is a warning message"),
    map[string]interface{}{"key": "value"},
)

/**
{
    "code": 1002,
    "message": {
        "content": "this is a warning message",
        "type": "warning"
    },
    "data": {
        "key": "value"
    }
}
*/
```

#### `RspDataPag(code interface{}, message Message, data interface{}, pag Pagination)`

```go
ogs.RspDataPag(
    0,
    ogs.InfoMsg("this is a info message"),
    []string{"1", "2"},
    ogs.NewPag(1, 10, 2),
)

/**
{
    "code": 0,
    "message": {
        "content": "this is a info message",
        "type": "info"
    },
    "data": [
        "1",
        "2"
    ],
    "pagination": {
        "current_page": 1,
        "total_pages": 5,
        "total_count": 10,
        "per_page": 2
    }
}
*/
```

-----
### You can change +CodeOK+ to any value

```go
ogs.ChangeOKCode("200")
```

-----

### other methods
```go
RspError(code interface{}, msgContent string)
RspDataOK(msgContent string, data interface{})
RspDataPagOK(msgContent string, data interface{}, pag Pagination)
```