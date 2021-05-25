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
	"encoding/json"
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


		//
		reminder,_ := json.Marshal(args.Reminder)

		// 插入数据库
		demoModel.Insert(map[string]interface{}{
			demoModel.Field.F_reminder:reminder,
		})


	return nil
}

// 获取信息
func (d *DemoLogic)GetDem(ctx context.Context) error{

	demoModel := new(models.DemoModel).Init()

	// 查找数据
	demoInfo := demoModel.Find(map[string]interface{}{

	})

	//map转struct  mapstructure.WeakDecode

	var reminder  []demo.ReminderInfo
	_ = json.Unmarshal([]byte(demoInfo[demoModel.Field.F_reminder].(string)),&reminder)

	// 图片信息处理




	return nil
}


// 编辑信息
func (d *DemoLogic)EditDemo(ctx context.Context, args * demo.ArgsDemo) error {

	// 先查询信息是否存在


	// 验证数据是否存在


	//对数据进行更新


	return nil
}

//获取列表
func (d *DemoLogic)ListDemo(ctx context.Context, demoId int,  ) error{


	return nil
}