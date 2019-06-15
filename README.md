## Installation
go-koa requires Golang v1.11.0 or higher for go.mod support.

```go
import "github.com/EvanLiu2968/go-koa"
```
## Hello Koa
```go
package main

import "github.com/EvanLiu2968/go-koa"

func main() {
  app := koa.New()
  // response
  app.Use(func(ctx koa.Context){
    ctx.SetBody("hello, koa")
  })

  app.Listen(3000)
}
```
## Context
主要方法都集中在ctx
```go
type Context interface {
  SetBody(body string)
  // 获取响应头
  GetContentType()
  // 设置响应头
  SetContentType(cType string)
  // 获取header
  GetHeader(name string)
  // 设置header
  SetHeader(name string, value string)
  // 获取cookie
  GetCookie(name string)
  // 设置cookie
  SetCookie(name string, value string)
  // ...
}
```

## License
MIT