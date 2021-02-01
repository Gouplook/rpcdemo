/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/27 13:32
@Description:

*********************************************/
package lib

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"github.com/wenzhenxi/gorsa"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"rpcdemo/rpcinterface/interface/pay"
	"rpcdemo/upgin/logs"
	"sort"
	"strings"
)

type icbcpay struct {

	// 公共类型参数
	host       string
	merId      string // 商户编号
	merPrtclNo string // 收单产品协议编号
	notifyUrl  string // 异步通知商户URL

	// 签名加密所需参数
	apigwPublicKey string // 网关公钥
	pre            string // 签名前缀
	singType       string // 签名类型
	encryptType    string // 加密方式

	privateKey string // 私钥匙
	encryptKey string // 加密key
}

const (
	allchar = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

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
	// 获取URL urlValue 中含有请求数据
	_, urlValue, _, err := i.URLValues(bizContentMap)
	urlStr := urlValue.String()
	return urlStr, err
}

// 微信小程序支付 (采用AES方式进行加密）
// 微信公众好支付（参数不同, 在payInfo和bizContent)配置参数
func (i *icbcpay) PayWxOfficial(payInfo *pay.PayInfo) (*gjson.Result, error) {
	realAmount, _ := decimal.NewFromString(payInfo.RealAmount)
	amount := realAmount.Mul(decimal.New(100, 0)).String()
	bizContextMap := map[string]string{
		"mer_id":       i.merId,         // 必须 商户编号
		"mer_prtcl_no": i.merPrtclNo,    // 必须 收单产品协议编号
		"order_id":     payInfo.OrderSn, // 商户订单号
		"amount":       amount,          // 必须 订单金额，单位为分
		"mer_url":      i.notifyUrl,     // 异步通知商户URL，端口必须为443或80
	}

	// 1：拼接请求参数及URL （URLValue)
	urlStr, _, pValue, err := i.URLValues(bizContextMap)
	if err != nil {
		return nil, nil
	}

	// 2：请求操作（doRequest)
	bytes, err := i.DoRequest("POST", urlStr, pValue)
	if err != nil {
		return nil, nil
	}

	// 3：请求数据进行解析并获取响应参数
	// response_biz_content：响应参数集合,包含公共和业务参数
	resultBody := gjson.Parse(string(bytes))
	respBizContentResult := resultBody.Get("response_biz_content")
	respBizContext := respBizContentResult.String()

	// 4：对数据进行解密
	decipherBizContent, err := i.AesDecipherBizContent(respBizContext, i.encryptKey)
	if err != nil {
		return nil, nil
	}

	// 5: 对解密后的数据进行验签
	sign := resultBody.Get("sign").String()
	err = i.veriSign(decipherBizContent, sign, "RSA")
	if err != nil {
		return nil, nil
	}
	returnCode := respBizContentResult.Get("return_code").Int()
	if returnCode != 0 {
		returnMsg := respBizContentResult.Get("return_msg").String()
		logs.Info("PayWxOfficial = %d, 错误信息：= %s", returnCode, returnMsg)
		return nil, nil
	}

	// 6: 返回支付信息，微信还是支付宝支付
	res := respBizContentResult.Get("wx_data_package") //微信支付数据包

	return &res, nil

}

// APP 支付
func (i *icbcpay)PayApp(payInfo *pay.PayInfo)(*gjson.Result, error){
	realAmount, _ := decimal.NewFromString(payInfo.RealAmount)
	amount := realAmount.Mul(decimal.New(100,0)).String()
	bizContent := map[string]string {
		"mer_id":       i.merId,         // 必须 商户编号
		"mer_prtcl_no": i.merPrtclNo,    // 必须 收单产品协议编号
		"order_id":     payInfo.OrderSn, // 商户订单号
		"amount":       amount,          // 必须 订单金额，单位为分
		"mer_url":      i.notifyUrl,     //异步通知商户URL，端口必须为443或80
	}
	// 获取URL
	urlStr, _,pValue ,err := i.URLValues(bizContent)
	if err != nil {
		return nil, nil
	}

	bytes, err := i.DoRequest("POST",urlStr,pValue)

	//
	respBody := gjson.Parse(string(bytes))
	responBizContent := respBody.Get("response_biz_content")
	responBizContentStr := responBizContent.String()

	// 对获取的字符串，进行解密
	decipherBizContent, err := i.AesDecipherBizContent(responBizContentStr,i.encryptKey)

	resultCode := respBody.Get("return_code").Int()
	resultMsg := respBody.Get("return_msg").String()
	if  resultCode != 0 {
		logs.Info("PayApp err : %d,%s ",resultCode, resultMsg)
		return nil, nil
	}
	// 	解密后验签
	sign := respBody.Get("sign").String()
	err = i.veriSign(decipherBizContent, sign,"RSA")
	if err != nil {
		logs.Info("veriSign failed")
		return nil, nil
	}
	// 判断支付渠道
	// 1= 微信  2 = 支付宝
	var res gjson.Result
	if payInfo.PayChannel == 1 {
		res = responBizContent.Get("wx_data_package")
	}else if payInfo.PayChannel == 2 {
		res = responBizContent.Get("wx_data_package")
	}

	return &res,nil
}

// 加密	 公钥加密
// @bizContent			明文（请求数据）
// @encryptBizContent 	密文
// @encryptKey  		加密key
// 原理： 明文 + 加密key = 密文
func (i *icbcpay) AesEncryptBizContent(bizContent []byte) (encryptBizContent string, encryptKey string, err error) {
	// 检测签名类型和加密方式
	if i.singType == "CA" && i.encryptType == "AES" {
		// 生成AesKey
		datakey := i.createAesKey(16)
		// 	利用公钥加密
		encryptKey, err = gorsa.PublicEncrypt(datakey, i.apigwPublicKey)
		if err != nil {
			return
		}
		encryptBizContent = i.aesEncrypt(bizContent, []byte(datakey))
	}
	return
}

func (i *icbcpay) aesEncrypt(origData []byte, key []byte) string {
	newCiphers, _ := aes.NewCipher(i.generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	srcData := make([]byte, length*aes.BlockSize)
	copy(srcData, origData)
	pad := byte(len(srcData) - len(origData))
	for i := len(origData); i < len(srcData); i++ {
		srcData[i] = pad
	}
	dstData := make([]byte, len(srcData))
	// 分组分块加密
	for bs, be := 0, newCiphers.BlockSize(); bs <= len(origData); bs, be = bs+newCiphers.BlockSize(), be+newCiphers.BlockSize() {
		newCiphers.Encrypt(dstData[bs:be], srcData[bs:be])
	}
	return base64.StdEncoding.EncodeToString(dstData)
}

// 解密  利用私钥解密
// @encryptBizContent	:加密后的密文
// @encryptKey			:加密key
// @decipherBizContent	:解密后明文
func (i *icbcpay) AesDecipherBizContent(encryptBizContent string, encryptKey string) (decipherBizContent string, err error) {
	// 1：公钥加密，私要解密
	decry, err := gorsa.PriKeyDecrypt(encryptKey, i.privateKey)
	srcData, err := base64.StdEncoding.DecodeString(encryptBizContent)

	byteKey := []byte(decry)
	// 2：根据私钥解密后，创建一个新密码
	newCiphers, _ := aes.NewCipher(i.generateKey(byteKey))
	dstData := make([]byte, len(srcData))
	// 分组分块解密
	for bs, be := 0, newCiphers.BlockSize(); bs < len(srcData); bs, be = bs+newCiphers.BlockSize(), be+newCiphers.BlockSize() {
		newCiphers.Decrypt(dstData[bs:be], srcData[bs:be])
	}
	trim := 0
	if len(dstData) > 0 {
		trim = len(dstData) - int(dstData[len(dstData)-1])
	}
	return string(dstData[:trim]), nil
}

// 生成一个AesKey
func (i *icbcpay) createAesKey(size int) (res string) {
	for i := 0; i < size; i++ {
		res += string(allchar[rand.Intn(62)])
	}
	return
}

func (i *icbcpay) generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// @bizContent biz_content参数说明
// 功能： 1：组装公共参数
//       2：对组装的参数进行key排序
//       3：签名
//       4：获取URL
// 返回值
// @urlStr   URL 字符串
// @urlValue URL+ "?" + 请求数据
// @pValues  请求参数
func (i *icbcpay) URLValues(bizContext map[string]string) (urlStr string, urlValue *url.URL, pVaules url.Values, err error) {
	var p = url.Values{}
	// 公共参数
	p.Add("key", "value")
	// ...

	// 拼接：公共参数 + biz_context
	bytes, err := json.Marshal(bizContext)
	if err != nil {
		return
	}

	p.Add("biz_content", string(bytes))

	// url.values --> map[string][string]
	data := make(map[string]string)
	for key, _ := range p {
		data[key] = p.Get(key)
	}

	// 对参数进行排序
	dataStr := i.requestParameterOrder(data)
	// 这些参数放到成配置文件
	// 签名
	sign, err := i.signStr(i.pre, dataStr, i.singType, i.privateKey)
	p.Add("sign", sign)

	urlStr = i.host + i.pre
	pVaules = p

	// 解析
	urlValue, err = url.Parse(urlStr + "?" + p.Encode())
	if err != nil {
		return
	}
	return
}

//验签 公钥签名
func (i *icbcpay) veriSign(data, sign, signType string) (err error) {
	// 目前 验签签名只支持RSA方式
	switch signType {
	case "RSA":
		err = gorsa.VerifySignSha1WithRsa(data, sign, i.apigwPublicKey)
	case "RSA2":
		err = gorsa.VerifySignSha256WithRsa(data, sign, i.apigwPublicKey)
	default:
	}
	return
}

// 签名 私钥签名
// @pre         签名分两种类型，一种带有前缀字符串，一种不带前缀字符串
// @data        签名数据
// @singType    签名类型
// @privateKey  私钥
func (i *icbcpay) signStr(pre string, data string, singType string, privateKey string) (sign string, err error) {
	if len(pre) == 0 { // 不带前缀字符串签名
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

// 组合的参数排序
// @data：	=common+biz_context数据
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
// @urlValue：	 = url + 请求数据
// @method：	请求方法
// @url：		URL
// @data        请求数据
// 功能：		：数据请求
func (i *icbcpay) DoRequest(method string, url string, data url.Values) (bytes []byte, err error) {
	// 1:读取数据
	var buf io.Reader
	buf = strings.NewReader(data.Encode())
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return
	}
	// 2：设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	// 3：请求动作
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// 4: 关闭请求
	if resp != nil {
		defer resp.Body.Close()
	}
	// 5：读取Body体数据
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
