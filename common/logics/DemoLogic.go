/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  demo
 * @Version: 1.0.0
 * @Date: 2021/1/17 18:55
 */
package logics

import (
	"context"
	"rpcdemo/rpcinterface/interface/demo"
)

type DemoLogic struct {

}

func (d *DemoLogic)DemoSample(ctx context.Context, args *demo.ArgsDemo, reply *demo.ReplyDemo) error {

	return nil
}
