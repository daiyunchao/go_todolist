package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"go_todolist/common/define"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	listenAddress string
	r             *gin.Engine
}

// CreateHttpServer 创建一个服务器
func CreateHttpServer(listenAddress string) *HttpServer {
	httpServer := &HttpServer{
		listenAddress: listenAddress,
		r:             gin.Default(),
	}
	return httpServer
}

// Run 启动服务器
func (server *HttpServer) Run() {
	go server.r.Run(server.listenAddress)
}

// RegisterRoutes 注册该服务器的路由
func (server *HttpServer) RegisterRoutes(routeName string, handle func(ctx context.Context, request *define.Request) *define.Response) {
	server.r.TrustedPlatform = "X-Client-IP"
	server.r.POST(routeName, func(c *gin.Context) {
		reqId := time.Now().UnixNano()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "reqId", reqId)
		handleChan := make(chan *define.RetResponse, 1)
		go func() {
			//处理请求
			request, requestErr := server.handlerRequest(c)
			if requestErr.Error != "" {
				handleChan <- requestErr
				return
			}
			//执行主逻辑回调
			handlerRes := handle(ctx, &request)

			//处理返回
			ret := server.handlerResponse(reqId, handlerRes)
			handleChan <- ret
		}()
		select {
		case retRes := <-handleChan:
			c.JSON(200, retRes)
		case <-ctx.Done():
			retRes := define.RetResponse{
				Code:  define.Timeout,
				Error: define.TimeoutMsg,
			}
			c.JSON(200, retRes)
			fmt.Printf("reqId: %d,finish,res: %s\n", reqId, retRes)
		}
	})
}

func (server *HttpServer) RegisterStatic(routeName string, path string) {
	server.r.Static(routeName, path)
}

// Stop 服务器停止
func (server *HttpServer) Stop() {

}

func (server *HttpServer) handlerRequest(c *gin.Context) (define.Request, *define.RetResponse) {
	//解析出参数
	request := define.Request{
		Data: make(map[string]any),
	}
	module := c.Param("module")
	method := c.Param("method")
	request.Module = module
	request.Method = method
	params := define.OriJsonRequest{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		return request, CreateRetResponseError(define.ParamError, define.ParamErrorMsg)
	}
	request.Data = params.Data
	return request, &define.RetResponse{}
}

func (server *HttpServer) handlerResponse(reqId int64, res *define.Response) *define.RetResponse {
	retRes := define.RetResponse{}
	bytes, err := json.Marshal(res.Data)
	if err != nil {
		return CreateRetResponseError(define.ServerError, "服务器内部错误")
	}
	retRes.Code = res.Code
	retRes.Error = res.Error
	retRes.Data = string(bytes)
	return &retRes
}
