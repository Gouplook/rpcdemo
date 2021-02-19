/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/19 14:24
@Description:

*********************************************/
package service

import (
	"context"
	"rpcdemo/common/logics"
	"rpcdemo/rpcinterface/interface/demo"
)

type Demo struct {

}

// DmeoSample 对外或rpc远程调用的开放接口
func (d *Demo)DemoSample(ctx context.Context, args *demo.ArgsDemo, reply *demo.ReplyDemo) error{
	// 初始化common中的logic，并调用结构体的方法
	demoLogic := new(logics.DemoLogic)
	err := demoLogic.DemoSample(ctx, args, reply)
	if err != nil {
		return err
	}
	return nil
}