package main

import (
	"awesomeProject/rpc/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {

	rpc.RegisterName("MathService", new(server.MathService))

	rpc.HandleHTTP() //新增的

	l, e := net.Listen("tcp", ":1234")

	if e != nil {

		log.Fatal("listen error:", e)

	}

	http.Serve(l, nil) //换成http的服务
}
