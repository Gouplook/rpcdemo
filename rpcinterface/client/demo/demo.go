/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/19 13:59
@Description:

*********************************************/
package demo

import (
	"context"
	"rpcdemo/rpcinterface/client"
	"rpcdemo/rpcinterface/interface/demo"
)

type DemoClient struct {
	client.Baseclient
}

func(d *DemoClient)DemoSample(ctx context.Context, args *demo.ArgsDemo, reply *demo.ReplyDemo) error{
	// client 返回
	return d.Call(ctx, "DemoSample", args, reply)
}
