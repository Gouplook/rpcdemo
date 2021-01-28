/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/4 11:24
@Description:

*********************************************/
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestCreateModel(t *testing.T) {
	// 创建数据库Model
	// utils.CreateModel("bank_bus")

	str := "https://api.shutung.com/v1/pay/sandpay"


	s := strings.Index(str, "/")
	fmt.Println(s)

}

func TestRe(t *testing.T) {
	now := time.Now()
	commonData := map[string]string{
		"FORMAT":    "constkey.FORMAT_JSON",
		"CHARSET":   "constkey.CHARSET_UTF8",
		"APP_ID":    "i.appId",
		"MSG_ID":    "GetMsgId()",
		"SIGN_TYPE": "constkey.SIGN_TYPE_RSA2",
		"TIMESTAMP": now.Format("2006-01-02"),
	}
	bizContentMap := map[string]string{
		"mer_id":          "i.merId",                 // 必须 商户编号
		"mer_prtcl_no":    "i.merPrtclNo",            // 必须 收单产品协议编号
		"order_id":        "payInfo.OrderSn",         // 必须 商户订单号，只能是数字、大小写字母，且在同一个商户号下唯一
		//"order_date_time": getOrigDateTime(),       // 必须 交易日期时间，格式为yyyy-MM-dd'T'HH:mm:ss
		"order_date_time": "2020-02-02T10:20:20",       // 必须 交易日期时间，格式为yyyy-MM-dd'T'HH:mm:ss
		"amount":         " amount",                  // 必须 订单金额，单位为分
		"cur_type":        "constkey.FEE_TYPE_CNY",   // 必须 交易币种，目前工行只支持使用人民币（001）支付
		"notify_url":      "i.notifyUrl",             // 必须 异步通知商户URL，端口必须为443或80
		"icbc_appid":      "i.appId",                 // 必须 工行API平台的APPID
	}
	bytes, _ := json.Marshal(bizContentMap)
	commonData["biz_context"] = string(bytes)
	fmt.Println(commonData)

	//str := lib.IcbcPay.RequestParameterOrder(commonData)
	var p = url.Values{}
	p.Add("app_id", "this.appId")
	p.Add("biz_content",string(bytes))

	//fmt.Println("p",p)
	data :=make(map[string]string)
	for k := range p {
		var value = p.Get(k)
		data[k] = value

	}
	fmt.Println(data)






}