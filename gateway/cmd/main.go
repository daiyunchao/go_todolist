package main

import (
	"go_todolist/common/lib"
	"go_todolist/gateway/route"
)

func main() {
	//启动服务器
	app := &App{}
	app.startServer()
	select {}
}

type App struct {
}

func (app *App) startServer() {
	httpServer := lib.CreateHttpServer("127.0.0.1:8000")
	httpServer.Run()
	route := route.GetRouteInstance()
	httpServer.RegisterRoutes("/:module/:method", route.DispatchApiRequest)
}
