# ogs-go

Return a well-structured data format for golang api

### `Rsp(code interface{}, message message)`

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/xandercheung/ogs-go"
)

func main() {
	v := ogs.Rsp(0, ogs.NewMsg("this is a message", "success"))
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
### `RspData(code interface{}, message message, data interface{})`

```go
ogs.RspData(
    "1001",
    ogs.WarningMsg("this is a warning message"),
    map[string]interface{}{"key": "value"},
)
/**
{
    "code": "1001",
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

### `RspDataPag(code interface{}, message message, data interface{}, pag paginate)`

```go
ogs.RspDataPag(
    "0",
    ogs.InfoMsg("this is a info message"),
    []string{"1", "2"},
    ogs.NewPag(1, 10, 2),
)

/**
{
    "code": "0",
    "message": {
        "content": "this is a info message",
        "type": "info"
    },
    "data": [
        "1",
        "2"
    ],
    "paginate": {
        "current_page": 1,
        "total_pages": 5,
        "total_count": 10,
        "per_page": 2
    }
}
*/
```

-----

### other methods
```go
RspOK(msgContent string)
RspError(code interface{}, msgContent string)
RspDataOK(msgContent string, data interface{})
RspDataPagOK(msgContent string, data interface{}, pag paginate)
```