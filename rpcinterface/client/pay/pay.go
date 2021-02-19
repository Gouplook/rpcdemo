/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 16:04
@Description:

*********************************************/
package pay

import (
	"context"
	"rpcdemo/rpcinterface/client"
	"rpcdemo/rpcinterface/interface/pay"
)

type ICBCPay struct {
	client.Baseclient
}

func (i *ICBCPay) Init() *ICBCPay {
	i.ServiceName = "rpc_demo"
	i.ServicePath = "Pay"
	return i
}

// H5支付宝支付
func (i *ICBCPay)H5AliPay(ctx context.Context, args *pay.PayInfo, reply *string) error {
	return i.Call(ctx, "H5AliPay", args, reply)
}
