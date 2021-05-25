/**
 * @Author: yinjinlin
 * @File:  ImageModel
 * @Description:
 * @Date: 2021/5/25 下午1:27
 */

package models

type ImageModel struct {

	Field ImageModelField
}

//表字段
type ImageModelField struct {
	T_table       string `default:"image"`
	F_id          string `default:"id"`
	F_name        string `default:"name"`
	F_md5         string `default:"md5"`
	F_sha256      string `default:"sha256"`
	F_size        string `default:"size"`
	F_create_time string `default:"create_time"`
	F_hash        string `default:"hash"`
	F_path        string `default:"path"`
	F_ext         string `default:"ext"`
	F_type        string `default:"type"`  // 上传图片类型
}


func (i *ImageModel)Init() *ImageModel{
	// .....

	return nil
}

// 插入数据
func (i *ImageModel)Insert(data map[string]interface{}) int {
	//.....

	rs := 2
	return rs
}

// 查重
func (i *ImageModel) FindRepeatImage(md5 string, sha256 string, size int64) map[string]interface{} {
	//....
	rs := make(map[string]interface{})
	// rs := i.Model.Where(map[string]interface{}{
	// 	i.Field.F_md5:    md5,
	// 	i.Field.F_sha256: sha256,
	// 	i.Field.F_size:   size,
	// }).Find()

	return rs
}