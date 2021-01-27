/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 17:22
@Description:

*********************************************/
package pay

import (
	"rpcdemo/rpcinterface/interface/pay"
	"github.com/tidwall/gjson"
)

type Pay interface {
	// 二维码支付
	PayQr(info *pay.PayInfo) (string, error)
	// h5支付
	PayH5(info *pay.PayInfo) (string, error)
	// 微信支付
	PayWxApp(info *pay.PayInfo) (*gjson.Result, error)
	// 微信公众号支付
	PayWxOfficial(info *pay.PayInfo) (*gjson.Result, error)
	// app支付
	PayApp(info *pay.PayInfo) (*gjson.Result, error)
}



//  获取支付渠道
//  ICBC  CCBC  ABC...等银行
func GetPay(payChannel int) (*Pay, error){

	return nil ,nil
}
