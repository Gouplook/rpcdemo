/**
 * @Author: yinjinlin
 * @File:  elasticClient
 * @Description:
 * @Date: 2021/5/26 下午2:14
 */

package upelastic

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	maxSearchNum = 2000  // 单次搜索最大返回数量,防止陷入深度搜索，拖垮搜索服务器
	numShareds   = 3     // 每个索引的分片数
	numReplicas  = 1     // 每个索引每个分片对应的副本数
	indexPrefix  = "up-" // 索引前缀
	indexPre     string

	// 自定义分析器
	// standard=es默认分析器,对中文支持较差
	// ik_smart=插件，中文分析器,对中文分词支持优秀
	// ik_max_word=插件，加强版ik_smart,对短语分词更细腻
	analyzer = []string{
		"standard",
		"ik_smart",
		"ik_max_word",
	}

	client *elasticsearch.Client
)

type ElasticClient struct {
	index     string                              // 索引名称
	indexType string                              // 索引类型
	term      []map[string]map[string]interface{} // 聚合筛选条件
	terms     []map[string]map[string]interface{} // 聚合筛选条件In
	mustNot   []map[string]map[string]interface{} // 聚合排除筛选条件
	should    []map[string]map[string]interface{} // 聚合筛选条件Or
	ranges    []map[string]map[string]interface{} // 范围筛选条件
	keyword   []map[string]map[string]interface{} // 文本筛选
	sort      []map[string]interface{}            // 排序
	limit     map[string]int                      // 分页
	geo       map[string]interface{}              // 地理位置
	factor    map[string]interface{}              // 欢迎度
	lastQuery string                              // 打印请求json

	// ES request Client
	requestClient *elasticsearch.Client
	// ES response Client
	responseClient *esapi.Response
}

func init() {
	// ES 配置参数读取
	hostAdd := ""
	indexPre = ""
	username := ""
	password := ""

	var err error
	if hostAdd == "" {
		err = errors.New("es hostaddr has no config")
		log.Panicln(err)
	}

	// ES 设置config配置
	cfg := elasticsearch.Config{
		Addresses: []string{"xxx"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 5,
			Proxy: func(req *http.Request) (*url.URL, error) {
				if username != "" && password != "" {
					req.SetBasicAuth(username, password)
				}
				return nil, nil
			},
			DialContext: (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}

	// 初始化Client
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		err = errors.New(fmt.Sprintf("es NewCient error,error is %s", err.Error()))
	}

}

// 新增/更新一个文档
// @param  interface{} id 文档ID
// @param  map[string]interface{}   文档数据
func (e *ElasticClient) PostDoc(id interface{}, data map[string]interface{}) bool {

	// 如果为整型则转字符串类型
	if vs, ok := id.(int); ok {
		id = strconv.Itoa(vs)
	}

	idF := id.(string)
	jsonBody, _ := json.Marshal(data)
	req := esapi.IndexRequest{
		Index:        e.index,
		DocumentType: e.indexType,
		DocumentID:   idF,
		Body:         bytes.NewReader(jsonBody),
	}

	res, err := req.Do(context.Background(), e.requestClient)
	if err != nil {
		fmt.Println("req.Do Error,Error is ", err.Error())
		return false
	}
	e.responseClient = res

	defer e.close()
	var maps = make(map[string]interface{})
	responseStr := res.String()[strings.Index(res.String(), "{"):]

	err = json.Unmarshal([]byte(responseStr), &maps)
	if maps["result"] != "" {
		return true
	} else {
		fmt.Println("req.Do Error,Error is", maps["error"])
		return false
	}
}

// 更新指定文档字段值
// @param string docId           文档id
// @data map[string]interface    更新内容
// map[string]interface{}{"fieldeName":"value","fieldName1":[]interface{"+",1}} 可自增自减
func (e *ElasticClient) UpdateById(docId string, data map[string]interface{}) bool {
	source := ""
	params := make(map[string]interface{})

	//
	for i, v := range data {
		if k, ok := v.([]interface{}); ok {
			if len(k) == 2 {
				e.arrayData(i, k[0].(string), k[1], &source, &params)
			}
		} else if k, ok := v.([]string); ok {
			if len(k) == 2 {
				e.arrayData(i, k[0], k[1], &source, &params)
			}
		} else {
			source += "ctx._source." + i + "=" + "params." + i + ";"
			params[i] = v
		}
	}
	body := map[string]interface{}{
		"script": map[string]interface{}{
			"source": source,
			"params": params,
		},
	}
	jsonBody, _ := json.Marshal(body)
	//
	res, err := e.requestClient.Update(e.index, docId, bytes.NewReader(jsonBody))
	defer e.close()
	if err != nil {
		fmt.Println("req.Do Error,Error is ", err.Error())
		return false
	}
	e.responseClient = res
	var maps = make(map[string]interface{})
	responseStr := res.String()[strings.Index(res.String(), "{"):]
	err = json.Unmarshal([]byte(responseStr), &maps)
	if _, ok := maps["error"]; ok {
		fmt.Println("req.Do Error,Error is", maps["error"])
		return false
	} else {
		return true
	}
}

// 根据过滤器条件更新文档
// @data map[string]interface    更新内容
// map[string]interface{}{"fieldeName":"value","fieldName1":[]interface{"+",1}} 可自增自减
func (e *ElasticClient) Update(data map[string]interface{}) bool {

	return true
}

// 删除文档
// @param  interface{} id 文档ID
// @return bool
func (e *ElasticClient) DeleteDoc(id interface{}) bool {

	return true
}

// 重置索引文档某个字段默认值
// @param  string      docField    字段名称
// @param  interface{} value       设置值
func (e *ElasticClient) ResetFileValue(docField string, value interface{}) bool {

	return true
}

// 文档字段值精确值筛选
// @param  string      docField 字段名
// @param  interface{} value    字段值 类型可以是整型、字符串
func (e *ElasticClient) SetFilter(docField string, value interface{}) *ElasticClient {

	return nil
}

// 文档字段值数组精确值筛选
// @param  string      docField 字段名
// @param  interface{} value    数组,查找多个精确值，使用数组形式 []int or []string
func (e *ElasticClient) SetFilterArray(docField string, value interface{}) *ElasticClient {

	return nil
}

// 文档字段值数组包含过滤
// @param  string      docField 字段名
// @param  interface{} value    数组,查找多个包含值，使用数组形式 []int or []string
func (e *ElasticClient) SetFilterIn(docField string, value interface{}) *ElasticClient {

	return nil
}

// 文档字段值精确值排除筛选
// @param  string      docField 字段名
// @param  int         value    字段值
func (e *ElasticClient) SetFilterFalse(docField string, value int) *ElasticClient {

	return nil
}

// 文档字段筛选 大于
func (e *ElasticClient) SetFilterGt(docField string, value interface{}) *ElasticClient {

	return nil
}

// 文档字段筛选 大于等于
func (e *ElasticClient) SetFilterGte(docField string, value interface{}) *ElasticClient {

	return nil
}

// 文档字段筛选 小于
func (e *ElasticClient) SetFilterLt(docField string, value interface{}) *ElasticClient {

	return nil
}

// 文档字段筛选 小于等于
func (e *ElasticClient) SetFilterLte(docField string, value interface{}) *ElasticClient {

	return nil
}

// 文档字段值范围筛选
// @param  string  docField            文档字段名称
// @param  interface{} minValue        字段值范围开始
// @param  interface{}  maxValue       字段值范围结束
// @param  ...interface{} params       可变参数 左边界【gt:大于、gte:大于等于】 右边界【lt:小于、lte:小于等于】
func (e *ElasticClient) SetFilterRange(docField string, minValue interface{}, maxValue interface{}, params ...interface{}) *ElasticClient {

	return nil
}

// 文档字段值范围筛选不在范围内
// @param  string  docField            文档字段名称
// @param  interface{} minValue        字段值范围开始
// @param  interface{}  maxValue       字段值范围结束
// @param  ...interface{} params       可变参数 左边界【gt:大于、gte:大于等于】 右边界【lt:小于、lte:小于等于】
func (e *ElasticClient) SetFilterNotRange(docField string, minValue interface{}, maxValue interface{}, params ...interface{}) *ElasticClient {

	return nil
}

// 设置排序方式
// @param  string     docField 文档排序字段
// @param  ...string  sort 排序模式 【desc:降序、asc:升序 默认=desc】
func (e *ElasticClient) SetSortMode(docField string, sort ...string) *ElasticClient {

	return nil
}

// 设置离我最近排序
// @param string    string       索引经纬度字段名
// @param float64   longitude    经度
// @param float64   latitude     纬度
// @param ...string params  可变参数、接收一个参数，排序单位，默认=m
func (e *ElasticClient) SetNearMode(locationField string, longitude float64, latitude float64, params ...string) *ElasticClient {

	return nil
}

// 根据经纬度获取附近指定公里数的矩形四点坐标
// @param float64 longitude    经度
// @param float64 latitude     纬度
// @param float64 distance     公里数
func (e *ElasticClient) SquarePoint(longitude float64, latitude float64) map[string]map[string]float64 {

	return nil
}

// 设置用户定位经纬度以及附近多少公里
// @param string   string     索引经纬度字段名
// @param float64    longitude    经度
// @param float64    latitude     纬度
func (e *ElasticClient) SetDistance(locationField string, longitude float64, latitude float64) *ElasticClient {

	return nil
}

// 设置搜索结果受欢迎程度排序,如调用该方法，则e.SetSortMode()方法就不需要用了
// @param  string docField     影响受欢迎排序字段 如点赞数字段、阅读量字段
// @param  ...string params    可变参数,包含(modifier,factor,boostMode)一种融入受欢迎度更好方式用 modifier 平滑 docField 的值
// 介绍参考：
// https://www.elastic.co/guide/cn/elasticsearch/guide/current/boosting-by-popularity.html
func (e *ElasticClient) SetFactor(docField string, params ...string) *ElasticClient {

	return nil
}

// 设置分页
// @param  int start
// @param  int limit
func (e *ElasticClient) SetLimit(start int, limit int) *ElasticClient {

	return nil
}

// 查询
func (e *ElasticClient) Query() map[string]interface{} {

	return nil
}

// 打印请求json
func (e *ElasticClient) GetLastQuery() {

}

// 检查文档是否存在
func (e *ElasticClient) DocExist(docId string) bool {

	return true
}

// 嵌套文档，根据文档id新增嵌套字段数据
func (e *ElasticClient) AddNestedDataByDocId(docId string, nestedField string, data map[string]interface{}) bool {

	return true
}

// ----?--
func (e *ElasticClient) arrayData(name string, condition string, value interface{}, source *string, params *map[string]interface{}) {

	switch strings.ToLower(condition) {
	case "inc":
	case "+":
		*source += "ctx_source" + name + "=ctx._source." + name + "+params." + name + ";"
		break
	case "dec":
	case "-":
		*source += "ctx_source" + name + "=ctx._source." + name + "-params." + name + ";"
		break
	}
	(*params)[name] = value
}

func (e *ElasticClient) close() {
	err := e.responseClient.Body.Close()
	if err != nil {
		fmt.Println("close Error,Error is ", err.Error())
	}
}
