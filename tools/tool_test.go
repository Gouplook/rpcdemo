/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 10:48
@Description:

*********************************************/
package tools

import (
	"fmt"
	"testing"
)


// 测试获取经纬度
// 先根据区域的Id进行数据库查找。获取地址，在根据地址信息获取经纬度。
func TestGetGaoByLatLng(t *testing.T) {
	//"上海浦东区156号","上海"
	mp,err := GetGaoByLatLng("浦东区156号","上海")
	if err != nil {
		return
	}
	t.Log(mp)
}

func TestArrayUniqueString(t *testing.T) {
	array := []string{
		"abc","bcd","wer","bcd",
	}

	fmt.Println("array =",array)
	temp := ArrayUniqueString(array)
	fmt.Println(temp)
}