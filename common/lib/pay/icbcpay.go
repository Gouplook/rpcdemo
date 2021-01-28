/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 17:39
@Description:

*********************************************/
package pay

import (
	"github.com/tidwall/gjson"
	"rpcdemo/common/lib"
	"rpcdemo/rpcinterface/interface/pay"
)

type icbc struct {

}

//生成二维码
func (i *icbc) PayQr(payInfo *pay.PayInfo) (string, error) {
	return "", nil
}

// H5支付
func (i *icbc) PayH5(payInfo *pay.PayInfo) (string, error) {
	return lib.IcbcPay.H5AliPay(payInfo)
}

// 微信小程序支付
func (i *icbc) PayWxApp(payInfo *pay.PayInfo) (*gjson.Result, error) {
	return lib.IcbcPay.PayWxOfficial(payInfo)
}
// 微信公共号支付
func (i *icbc) PayWxOfficial(info *pay.PayInfo) (*gjson.Result, error) {
	return nil, nil
}

// 微信App支付
//获取App支付信息
func (i *icbc) PayApp(payInfo *pay.PayInfo) (*gjson.Result, error) {
	return nil, nil
}

