package main

import (
	"fmt"
	"../lib"
)

func main() {
  app := koa.App()
  // response
  app.use(func(ctx koa.Context){
    ctx.body = "hello, koa"
  })

  app.listen(3000)
}
