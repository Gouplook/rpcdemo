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
	"rpcdemo/common/models"
	"rpcdemo/rpcinterface/interface/demo"
)

type DemoLogic struct {

}

func (d *DemoLogic)DemoSample(ctx context.Context, args *demo.ArgsDemo, reply *demo.ReplyDemo) error {
		demoModel := new(models.DemoModel).Init()
		wh := map[string]interface{}{}
		start,limit := args.GetStart(),args.GetPageSize()
		demoModel.SelectByPage(wh,start,limit)

	return nil
}
