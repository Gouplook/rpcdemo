/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  timetool
 * @Version: 1.0.0
 * @Date: 2021/1/17 19:14
 */
package tools

import (
	_const "rpcdemo/lang/const"
	"rpcdemo/upbase/common/toolLib"
	"strconv"
	"time"
)

// 字符串时间格式转time.Time
func ParseDateTime(str string, layout ...string)(t time.Time, err error){
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var format string
	if len(layout) == 0 {
		format = "2006-01-02 15:04"
	} else {
		format = layout[0]
	}
	t, err = time.ParseInLocation(format, str,loc)
	if err != nil {
		err = toolLib.CreateKcErr(_const.TIME_ERR)
		return
	}
	return
}

// 时间戳转换为dataTime  时间戳转化为字符串
//@param  int64  timestamp  时间戳
//@return string
func Timestamp2DataTime(timestampStr string, layout ...string) (dataTime string){
	timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)
	var format string
	if len(layout) == 0 {
		format = "2006-01-02 15:04:05"
	} else {
		format = layout[0]
	}
	dataTime = time.Unix(timestamp, 0).Format(format)
	return
}


// 2020-07-20 00:00:00~2020-07-20 23:59:59 时间范围
func TimeRange(now time.Time) (bTime, eTime time.Time) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	bTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc) // 2020-07-20 00:00:00
	eTime = bTime.AddDate(0, 0, 1).Add(-1 * time.Second)                   // 2020-07-20 23:59:59
	return
}
