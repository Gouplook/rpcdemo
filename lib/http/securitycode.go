/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/8 13:23
@Description:

*********************************************/
package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smallnest/rpcx/log"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type SecurityCode struct {
}


// 	请求体参数
type RequestBody struct {
	field string `json:"field"` // 请求字段

}
// 响应结构体
type CommonResponseBody struct {
	field string  `json:"field"`  // 响应字段
}

// 获取URl
func (s *SecurityCode) getUrl() (url string) {
	url = "http.www.baidu.com"
	return url
}

// Content-Type ：application/json
// 请求远程URL地址获取信息
func (s *SecurityCode) doJson(requestBody *bytes.Buffer)(commonResponseBody *CommonResponseBody, err error) {
	// 为何需要转换成字符串
	requestBodyStr := fmt.Sprintf("%s", requestBody)
	requestBodyStr = strings.Replace(requestBodyStr, " ", "", -1)
	requestBodyStr = strings.Replace(requestBodyStr, "\n", "", -1)

	// 获取URL
	url := s.getUrl()
	// 1： 建立连接
	request, err := http.NewRequest("Get/POST/XXX", url, bytes.NewBufferString(requestBodyStr))
	//request, err := http.NewRequest("Get",url,requestBody)

	// 2 ：设置Header信息，若头部有信息，可以设置header信息
	request.Header = map[string][]string{
		"channels": []string{"xxx"},
	}
	request.Header.Set("Content-Type", "application/json;charset='utf-8")

	// 3：获取返回信息
	client := &http.Client{}
	response, err := client.Do(request)  // 远程请求
	if err != nil {
		panic(err)
	}
	// 5：关闭请求
	defer response.Body.Close()

	// 4：数据处理（数据返回）
	if response.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		log.Info("bosc response:" + string(body))
		err := json.Unmarshal(body, &commonResponseBody)
		if err == nil {
			return commonResponseBody, nil
		}
	}
	return commonResponseBody, errors.New("request error:" + response.Status + " URL:" + response.Request.URL.String())
}

// 请求远程数据URL Content-Type ：multipart/form-data
// multipart/form-data with this Writer's Boundary.
func (s *SecurityCode) doMultipart(body *RequestBody)(commonResponseBody string, err error) {
	requestBody := new(bytes.Buffer)
	// 将请求体对象 序列化Buffer格式，以便接受数据
	json.NewEncoder(requestBody).Encode(body)

	// 处理multipart/form-data
	w := multipart.NewWriter(requestBody)
	url := s.getUrl()
	// 1：建立连接
	request, err := http.NewRequest("Get/POST/XXX", url, requestBody)
	// 2： 设置header信息
	request.Header.Set("Content-Type",w.FormDataContentType())
	// 3：请求响应动作
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	// 5：关闭连接
	defer response.Body.Close()
	// 4：数据返回
	if response.StatusCode == http.StatusOK{
		body ,_  := ioutil.ReadAll(response.Body)
		bodyString := string(body)
		return bodyString,nil
	}

	return commonResponseBody, errors.New("request error:" + response.Status + " URL:" + response.Request.URL.String())
}


// 获取远程端的接口数据
func (s *SecurityCode) GetSecurityCode(body *RequestBody) (commonResponseBody *CommonResponseBody, err error) {
	// Buffer 结构体
	requestBuffer := new(bytes.Buffer)
	// 将请求体对象 序列化 Buffer格式
	json.NewEncoder(requestBuffer).Encode(body)
	return s.doJson(requestBuffer)
}
