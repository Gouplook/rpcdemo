/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/8 15:36
@Description:

*********************************************/
package tools

import (
	"archive/zip"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chenhg5/collection"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)


//高德api，通过地址获取经纬度
//@param string         address 地址
//@param ...interface{} 市区(如：北京)
func GetGeoByGordes(address string, param ...interface{}) (map[string]interface{}, error) {
	//url := kcgin.AppConfig.String("gordes.url") + "?address=%s&city=%s&output=JSON&key=%s"
	//key := kcgin.AppConfig.String("gordes.key")

	url := "xxx"
	key := "YYY"

	var city string
	if len(param) > 0 {
		city = param[0].(string)
	}
	reqUrl := fmt.Sprintf(url, address, city, key)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var maps map[string]interface{}
	json.Unmarshal(body, &maps)
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

//获取两个经纬度之间的距离
func GetDistance(lng1, lat1, lng2, lat2 float64, param ...string) float64 {
	var unit = "m"
	if len(param) > 0 {
		if param[0] == "km" {
			unit = "km"
		}
	}
	radius := 6371000.0 //单位：米
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	distance := dist * radius
	if unit == "km" {
		distance = distance / 1000
	}
	return distance
}

//数组map排序
func SortsMap(field string, maps []map[string]interface{}) []map[string]interface{} {
	var mapData = make(map[string]interface{})
	var keys = make([]string, 0)
	for _, v := range maps {
		vs := v[field]
		if vp, ok := vs.(float64); ok {
			vs = strconv.FormatFloat(vp, 'f', -1, 64)
		}
		if vp, ok := vs.(int); ok {
			vs = strconv.FormatInt(int64(vp), 10)
		}
		if vp, ok := vs.(string); ok {
			vs = vp
		}
		mapData[vs.(string)] = v
		keys = append(keys, vs.(string))
	}
	sort.Strings(keys)
	remapData := make([]map[string]interface{}, 0)
	for _, v := range keys {
		remapData = append(remapData, mapData[v].(map[string]interface{}))
	}
	return remapData
}

//检测两个数组被包含关系
func CheckArrExsits(arr []string, checkArr []string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range checkArr {
		if !collection.Collect(arr).Contains(v) {
			return false
		}
	}
	return true
}



//Interface2String Interface2String
func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	}
	return ""
}

//Interface2Int Interface2Int
func Interface2Int(inter interface{}) int {
	switch inter.(type) {
	case int:
		return inter.(int)
	}
	return 0
}

//Interface2Int64 Interface2Int64
func Interface2Int64(inter interface{}) int64 {
	switch inter.(type) {
	case int64:
		return inter.(int64)
	}
	return 0
}

//Interface2Float64 Interface2Float64
func Interface2Float64(inter interface{}) float64 {
	switch inter.(type) {
	case float64:
		return inter.(float64)
	}
	return 0
}

//JSON2Map JSON2Map
func JSON2Map(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

//MD5String MD5String
func MD5String(str string) (md5Str string) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return
	}
	arr := m.Sum(nil)
	md5Str = fmt.Sprintf("%x", arr)
	return
}

// FormatTimeStr 将时间格式化成字符串
func FormatTimeStr(times int64) string {
	return time.Unix(times, 0).Format("2006-01-02")
}
func Timestamp2DateTime(timestampStr string, layout ...string) (dataTime string) {
	timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)
	var format string
	if len(layout) == 0 {
		format = "2006-01-02"
	} else {
		format = layout[0]
	}
	dataTime = time.Unix(timestamp, 0).Format(format)
	return
}

//GetSHA256FromFile 文件的 SAH-256 校验
func GetSHA256FromFile(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum, nil
}

//GetMd5FromFile 文件的 md5 校验
func GetMd5FromFile(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil

}

//压缩zip
func CompressZip(path string , dest string)(err error) {
	var fileObjs []*os.File
	files, _ := ioutil.ReadDir(path)
	for _, fi := range files {
		fileObj, err :=  os.Open(path + fi.Name())
		if err != nil{
			return err
		}
		defer fileObj.Close()
		fileObjs = append(fileObjs,fileObj)
	}
	d, err := os.Create(dest)
	if err != nil{
		return err
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()

	for _, file := range fileObjs {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}


//获取格式化私钥
func FormatPrivateKey(key string ) string {
	var publicHeader = "-----BEGIN RSA PRIVATE KEY-----\n"
	var publicTail = "-----END RSA PRIVATE KEY-----\n"
	var temp string
	split(key,&temp)
	return publicHeader+temp+publicTail
}

//64个字符换行
func split(key string,temp *string){
	if len(key)<=64 {
		*temp = *temp+key+"\n"
	}
	for i:=0;i<len(key);i++{
		if (i+1)%64==0{
			*temp = *temp+key[:i+1]+"\n"
			key = key[i+1:]
			split(key,temp)
			break
		}
	}
}

// 结构体字符串字段去除空格换行
func TrimStruct(inStructPtr interface{})(err error)  {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		err = errors.New("必须传入指针")
	}
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		if t.Type.Kind() == reflect.String{
			v := f.String()
			v = strings.Replace(v, " ", "", -1)
			v = strings.Replace(v, "\n", "", -1)
			v = strings.Replace(v, "\r", "", -1)
			f.Set(reflect.ValueOf(v))
		}
	}
	return
}