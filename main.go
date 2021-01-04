/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/4 10:22
@Description:

*********************************************/
package main

import (
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
	"rpcdemo/routers"
)

func main(){
	// 1: 打印环境变量


	// 2: 启动服务
	rpcServer := server.NewServer()
	routers.InitRpcRouters(rpcServer)
	// 需要地址, 地址暂时写死的
	address := "xxxx"


	// 3: 启动链路追踪(其他中间件）



	// 4：连接服务
	err := rpcServer.Serve("tcp", address)
	if err != nil {
		// rpc启动失败
		log.Info("failed to rpcserve:%v",err)
	}

}
