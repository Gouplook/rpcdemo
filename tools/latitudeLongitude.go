/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 10:25
@Description: 高德api，通过地址获取经纬度

*********************************************/
package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
//高德api，通过地址获取经纬度
//@param string         address 地址
//@param ...interface{} 市区(如：北京)
func GetGaoByLatLng(address string, param ...interface{}) (map[string]interface{}, error) {
	// 1:读取配置高德配置信息（高德提供）
	//url := upgin.AppConfig.String("gordes.url") + "?address=%s&city=%s&output=JSON&key=%s"
	//key := upgin.AppConfig.String("gordes.key")
	url := "https://restapi.amap.com/v3/geocode/geo" + "?address=%s&city=%s&output=JSON&key=%s"
	key := "af8f6a94db276c64e876ff3f82783414"
	var city string
	if len(param) > 0 {
		city = param[0].(string)
	}
	// 拼接请求URL
	reqUrl := fmt.Sprintf(url, address, city, key)
	// 2:请求高德api数据
	resp, err := http.Get(reqUrl)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	// 3:读取请求数据
	body, _ := ioutil.ReadAll(resp.Body)
	// 4: 数据序列化
	var maps map[string]interface{}
	_ = json.Unmarshal(body, &maps)
	// 5: 提取经纬度数据
	var data = make(map[string]interface{})
	if v, ok := maps["geocodes"]; ok {
		vs := v.([]interface{})
		if len(vs) > 0 {
			vs1 := vs[0].(map[string]interface{})
			s := strings.Split(vs1["location"].(string), ",")
			data["lng"] = s[0]
			data["lat"] = s[1]
		}
	}

	return data, nil
}
