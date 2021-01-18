/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 17:22
@Description:

*********************************************/
package pay

import (
	_const "rpcdemo/lang/const"
	"rpcdemo/rpcinterface/interface/order"
	"rpcdemo/rpcinterface/interface/pay"
	"rpcdemo/upbase/common/toolLib"
)

type Pay interface {
	// H5支付宝支付
	H5AliPay(payInfo *pay.ICBcPayInfo,reply *pay.ReplyH5AliPay) error
	// 获取支付渠道类型
	GetType() int


}

func GetPay(payChannel int) (*Pay, error) {
	switch payChannel {
	//case order.PAY_TYPE_ALI:
	//	p := Pay(new(ccb))
	//	return &p, nil
	//case order.PAY_CHANNEL_sand:
	//	p := Pay(new(sandPay))
	//	return &p, nil
	case order.PAY_CHANNEL_icbc:
		p := Pay(new())
		return &p, nil
	}
	return nil, toolLib.CreateKcErr(_const.PAY_CHANNEL_ERROR)
}
