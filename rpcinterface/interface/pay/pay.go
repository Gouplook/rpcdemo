/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/18 15:49
@Description:

*********************************************/
package pay

import "context"

type Wx struct {
	AppId  string // 微信appid
	OpenId string // 微信openid
}

type PayInfo struct {
	OrderSn           string // 订单编号
	BusId             int    // 购买商家总店id
	RealAmount        string // 订单总金额
	InsureAmount      string // 保险费用
	RenewInsureAmount string // 续保费用
	PlatformAmount    string // 平台手续费
	BusAmount         string // 商户收取金额
	PayChannel        int    // 支付渠道
	InsuranceChannel  int    // 保险渠道
	ChosePayType      int    // 支付方式
	Wx

	Version      string
	CreateIP     string
	StoreID      string
	PayExtra     string
	AccsplitFlag string
	SignType     string

	FormUrl string // 成功后跳转连接
}

type Pay interface {
	// 获取支付二维码
	PayQr(ctx context.Context, args *PayInfo, reply *string) error
	// 获取H5支付连接
	PayH5(ctx context.Context, args *PayInfo, reply *string) error
	// 获取小程序支付数据
	PayWxapp(ctx context.Context, args *PayInfo, reply *string) error
	// 微信公众号支付数据
	PayWxOfficial(ctx context.Context, args *PayInfo, reply *string) error
}