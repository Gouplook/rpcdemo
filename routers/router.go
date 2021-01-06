/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/4 10:33
@Description:

*********************************************/
package routers

import (
	"github.com/smallnest/rpcx/server"
	"log"
	"rpcdemo/service"
)


func InitRpcRouters(rpcServer *server.Server){
	// 注册路由 或 RegisterName
	err := rpcServer.Register(new(service.Student),"")
	if err != nil{
		// 打印日志
		log.Fatalf("failed to register rpcRouter: %v", err)
	}

}


