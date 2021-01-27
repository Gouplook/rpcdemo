/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/27 13:32
@Description:

*********************************************/
package lib

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/wenzhenxi/gorsa"
	"io"
	"net/http"
	"net/url"
	"rpcdemo/rpcinterface/interface/pay"
	"sort"
	"strings"
)

type icbcpay struct {
	//
	host       string
	merId      string // 商户编号
	merPrtclNo string // 收单产品协议编号

	pre        string // 签名前缀
	singType   string // 签名类型
	privateKey string // 私钥匙
}

var (
	IcbcPay *icbcpay
)

func init() {
	IcbcPay.host = "https://apipcs3.dccnet.com.cn"
	IcbcPay.pre = ""



}

// 二维码支付
// @amount 金额
// @title 支付标题
// @body 支付内容
// @formUrl 前端跳转的链接
// @return string 二维码内容
func (i *icbcpay) QrCode(payinfo *pay.PayInfo) (string, error) {

	// 1: 公共数据

	// 2: biz_context

	// 3: biz_context 数据加密 (封装加密方法）

	// 4: 公共参数+biz_context 签名，签名之前，需要对参数进行排序(封装签名方法）

	// 5：签名的字符串加入到参数中

	// 6: 签名好的数据 请求（request）

	return "", nil
}

// H5下单支付宝支付
func (i *icbcpay) H5AliPay(payInfo *pay.PayInfo) (string, error) {
	realAmount, _ := decimal.NewFromString(payInfo.RealAmount)
	amount := realAmount.Mul(decimal.New(1, 2)).String()

	// biz_context
	bizContentMap := map[string]string{
		"mer_id":       i.merId,         // 必须 商户编号
		"mer_prtcl_no": i.merPrtclNo,    // 必须 收单产品协议编号
		"order_id":     payInfo.OrderSn, // 商户订单号
		"amount":       amount,          // 必须 订单金额，单位为分
	}
	// 获取URL
	url, err := i.URLValues(bizContentMap)
	urlStr := url.String()
	return urlStr, err
}
// 微信小程序支付
func (i *icbcpay)PayWxOfficial(payInfo *pay.PayInfo){
	//


}


// 加密方法



// 签名
// @pre         签名分两种类型，一种带有前缀字符串，一种不带前缀字符串
// @data        签名数据
// @singType    签名类型
// @privateKey  私钥
func (i *icbcpay) signStr(pre string, data string, singType string, privateKey string) (sign string, err error) {
	if len(pre) == 0 {
		data += ""
	} else {
		data += pre + "?"
	}
	switch singType {
	case "RSA":
		sign, err = gorsa.SignSha1WithRsa(data, privateKey)
	case "RSA2":
		sign, err = gorsa.SignSha256WithRsa(data, privateKey)
	default:
	}
	return
}

// @bizContent biz_content参数说明
// 功能： 1：组装公共参数
//       2：对组装的参数进行key排序
//       3：签名
//       4：获取URL
func (i *icbcpay) URLValues(bizContext map[string]string) (urlValue *url.URL, err error) {
	var p = url.Values{}
	// 公共参数
	p.Add("key", "value")

	// 拼接：公共参数 + biz_context
	bytes, err := json.Marshal(bizContext)
	if err != nil {
		return nil, err
	}

	p.Add("biz_content", string(bytes))

	// url.values --> map[string][string]
	data := make(map[string]string)
	for key, _ := range p {
		data[key] = p.Get(key)
	}

	// 对参数进行排序
	datastr := i.requestParameterOrder(data)
	// 这些参数放到成配置文件

	// 签名
	sign, err := i.signStr(i.pre, datastr, i.singType, i.privateKey)
	p.Add("sign", sign)
	// 解析
	urlValue, err = url.Parse(i.host + i.pre + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return
}

// 组合的参数排序，common+biz_context 数据
func (i *icbcpay) requestParameterOrder(data map[string]string) string {
	pList := make([]string, 0, 0)
	for key, value := range data {
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var rpoStr = strings.Join(pList, "&")
	return rpoStr
}

// 请求
// @urlValue = url + 请求数据
func (i *icbcpay) DoRequest(method string, url string ,data string) {
	var buf io.Reader

	// 1:读取数据
	buf = strings.NewReader(data)
	req, err := http.NewRequest(method,url, buf)
	// 2：设置请求头
	req.Header.Set("Content-Type","application/x-www-form-urlencoded;charset=utf-8")
	// 3：请求动作
	client := &http.Client{}
	resp, err := client.Do(req)
	// 4: 关闭请求
	if resp != nil {
		defer resp.Body.Close()
	}


}
