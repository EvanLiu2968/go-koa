package koa

import (
	"net/http",
	"fmt"
)

type Application struct {
	key string
	handlers []*Handler
}

type Handler struct {
	route string
	handle func
}

func (app *Application) Use(route string, handle func) {
	h := Handler{
		route: route,
		handle: handle,
	}
  app.handler.push(h)
}

func (app *Application) Listen(port int) {
	for i, h := range app.handlers {
		http.HandleFunc(h.route, h.handle)
	}
	addr := fmt.Sprintf("127.0.0.0:%d", port)
	server, err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("go server is running at %s", addr)
	}
	return server
}

func New() {
	var app = Application{
		key: "app_key",
		middleware: []
	}
	return app
}