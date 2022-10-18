package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"lch-tests/handler"
	//"lch-tests/kube"
	//rpcClient "lch-tests/rpc-clients"
)

const (
	addr = "192.168.2.331:31878"
)

func Main(ctx context.Context) error {
	// 初始化kube配置信息
	//kube.InitKube()

	// 初始化rpc连接
	InitConnections(addr)

	engine := gin.New()
	handler := &handler.Handler{}
	engine.GET("/home", handler.Home)
	return engine.Run()
}
