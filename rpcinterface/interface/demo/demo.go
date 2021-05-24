/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/19 13:55
@Description:

*********************************************/
package demo

import (
	"context"
	"rpcdemo/rpcinterface/interface/common"
)

// 入参数
type ArgsDemo struct {
	common.Paging // 分页
	Reminder      []ReminderInfo      // 温馨提示
}

// 返回参数
type ReplyDemo struct {
}

// 温馨提示
 type ReminderInfo struct {
	 ReminderName    string
	 ReminderContent string
 }


// Demo接口
type Demo interface {
	// Demo样例
	DemoSample(ctx context.Context, args *ArgsDemo, reply *ReplyDemo) error
}
