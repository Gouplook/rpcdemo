/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 16:29
@Description:

*********************************************/
package service

import (
	"rpcdemo/common/logics/pay"
	payinterface "rpcdemo/rpcinterface/interface/pay"
	"context"
)

type Pay struct {

}

func (p *Pay) H5AliPay(ctx context.Context, args *payinterface.ICBcPayInfo, reply *string) (err error){
	*reply, err = new(pay.ICBCFundAgentLogic).H5AliPay(args)
	return err
}