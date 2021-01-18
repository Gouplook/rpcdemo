/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 15:46
@Description:

*********************************************/
package pay

import (
	"fmt"
	"rpcdemo/constkey"
	"rpcdemo/rpcinterface/interface/order"
	"rpcdemo/rpcinterface/interface/pay"
	redis "rpcdemo/upredis"
)

type ICBCFundAgentLogic struct {

}


// 工行 H5平台支付宝支付方式
func (i *ICBCFundAgentLogic)H5AliPay(icbcpayInfo  *pay.ICBcPayInfo) (string, error){
	//
	res := i.getCache(icbcpayInfo, "H5AliPay")
	if res == nil {
		// 获取H5 支付宝支付渠道

	}



	return "", nil

}

func (i *ICBCFundAgentLogic) getCache(payInfo *pay.ICBcPayInfo, payType string) interface{} {
	// 工行渠道
	if payInfo.PayChannel == order.PAY_CHANNEL_icbc {
		if payInfo.ChosePayType == order.PAY_TYPE_ALI {
			payInfo.ChosePayType = 1 // H5 支付宝支付
		}
	}

	// 支付请求缓存redis
	key := fmt.Sprintf(constkey.PAY_CACHE,payInfo.OrderId)
	res, _ := redis.RedisGlobMgr.Get(key)

	return res
}
