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
	"fmt"
	"rpcdemo/common/models"
	"rpcdemo/rpcinterface/client/file"
	"rpcdemo/rpcinterface/interface/demo"
)

type DemoLogic struct {

}

func (d *DemoLogic)DemoSample(ctx context.Context, args *demo.ArgsDemo, reply *demo.ReplyDemo) error {
		// 初始化model
		demoModel := new(models.DemoModel).Init()
		wh := map[string]interface{}{}
		start,limit := args.GetStart(),args.GetPageSize()
		// 分页查找
		demoModel.SelectByPage(wh,start,limit)

		// rpc之间调用，调用rpcFile
		rpcFile := new(file.Upload).Init()
		fmt.Println(rpcFile)

	return nil
}
