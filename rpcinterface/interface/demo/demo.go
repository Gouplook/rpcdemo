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
	common.Paging                // 分页
	Reminder      []ReminderInfo // 温馨提示

	// 图片存储问题，图片以字符串的形式存放在数据库中
	Picture []string //  相册图片
	ImgHash string   //  封面图片hansh串
}

// 返回参数
type ReplyDemo struct {
}

// 温馨提示
type ReminderInfo struct {
	ReminderName    string
	ReminderContent string
}

// 返回基础数据
type DemoBase struct {
}

// 返回列表
type ListDemo struct {
	TotalNum int        // 总条数
	List     []DemoBase // 列表基本信息

}

// Demo接口
type Demo interface {
	// Demo样例
	DemoSample(ctx context.Context, args *ArgsDemo, reply *ReplyDemo) error
}
