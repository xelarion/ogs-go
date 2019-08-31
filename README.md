# go-rsp
Standard response format for golang

RspBaseData:
```
{
  message: {
    message: "",
    type: ""
  },
  code: 0,
  data: null
}
```

RspPagData:
```
{
  message: {
    message: "",
    type: ""
  },
  code: 0,
  data: null,
  paginate: {
    current_page: 0,
    total_pages: 0,
    total_count: 0,
    per_page: 0
  }
}
```