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
  app := koa.App()
  // response
  app.use(func(ctx koa.Context){
    ctx.body = "hello, koa"
  })

  app.listen(3000)
}
```

License
MIT