/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 17:39
@Description:

*********************************************/
package pay

import (
	"github.com/shopspring/decimal"
	"rpcdemo/rpcinterface/interface/pay"
)

type icbcPay struct {

}


// 工行H5支付宝支付
func (i *icbcPay) H5AliPay(payInfo *pay.ICBcPayInfo,reply *pay.ReplyH5AliPay) error{
	realAmout, _ := decimal.NewFromString(payInfo.RealAmount)


}
