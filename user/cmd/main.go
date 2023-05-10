package main

import (
	"flag"
	"fmt"
	"go_todolist/user/proto"
)

var port = "50051"

func main() {
	flag.StringVar(&port, "p", "50051", "端口号")
	flag.Parse()
	rpcServer := proto.Server{}
	address := fmt.Sprintf("127.0.0.1:%s", port)
	go rpcServer.RegisterServer(address)
	fmt.Printf("listen address %s\n", address)
	select {}
}
