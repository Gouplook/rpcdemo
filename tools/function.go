/**
 * @Author: yinjinlin
 * @File:  function
 * @Description:
 * @Date: 2021/5/25 上午9:38
 */

package tools

import "fmt"

// 数组去重
func ArrayUniqueString(array []string) []string {
	if len(array) == 0 {
		return array
	}

	temp := make([]string, 0)
	for i := 0; i < len(array); i++ {
		repeat := false
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				repeat = true
				break
			}
		}
		if array[i] == "" {
			continue
		}
		if repeat == false {
			temp = append(temp, array[i])
		}

	}

	return temp
}

// map中根据字段涮选字段



// 获取文件类型的hash字符串
func GetFileNameByHash(hash string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", hash[0:3], hash[3:6], hash[6:9], hash[9:12], hash[12:15], hash[15:18], hash[18:])
}

//
func GetFileName(name string) string{
	if len(name) > 64{
		return name[0:64]
	}
	return name
}


