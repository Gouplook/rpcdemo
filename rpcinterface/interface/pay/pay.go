/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 15:49
@Description:

*********************************************/
package pay

import "context"

type ICBcPayInfo struct {
	MerId         string // 商户编号
	MerPrtclNo    string // 收单产品协议编号
	OrderId       string // 商户订单号
	OrderDateTime string // 交易日期时间，格式为yyyy-MM-dd'T'HH:mm:ss
	Amount        string // 订单金额，单位为分
	CurType       string // 交易币种，目前工行只支持使用人民币（001）支付
	Body          string // 商品描述
	NotifyUrl     string // 异步通知商户URL，端口必须为443或80 *
	IcbcAppId     string // 工行API平台的APPID

	PayChannel    int    // 支付渠道
	ChosePayType  int    // 支付方式
	RealAmount    string // 订单总金额

}

type ReplyH5AliPay struct {
}

type ICBCPay interface {
	// H5支付宝支付
	H5AliPay(ctx context.Context, args *ICBcPayInfo, reply *string) error
}
