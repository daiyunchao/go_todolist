package route

import (
	"context"
	"go_todolist/common/define"
	"go_todolist/gateway/handler"
	"sync"
)

type Route struct {
}

var once sync.Once
var instance *Route

func GetRouteInstance() *Route {
	once.Do(func() {
		instance = &Route{}
	})
	return instance
}

func (route *Route) DispatchApiRequest(ctx context.Context, request *define.Request) *define.Response {
	gatewayHandler := handler.GetGateway()
	var res = &define.Response{}
	switch request.Module {
	case "user":
		switch request.Method {
		case "login":
			res = gatewayHandler.Login(ctx, request)
		}
	default:
		res.Code = define.PathError
		res.Error = define.PathErrorMsg
	}
	return res
}
