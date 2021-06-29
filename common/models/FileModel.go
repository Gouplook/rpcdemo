/**
 * @Author: yinjinlin
 * @File:  FileModel
 * @Description:
 * @Date: 2021/5/25 下午4:45
 */

package models

// 表结构体
type FileModel struct {
	// Model *base.Model
	Field FileModelField
}

// 表字段
type FileModelField struct {
	T_table       string `default:"file"`
	F_id          string `default:"id"`
	F_name        string `default:"name"`
	F_md5         string `default:"md5"`
	F_sha256      string `default:"sha256"`
	F_size        string `default:"size"`
	F_create_time string `default:"create_time"`
	F_hash        string `default:"hash"`
	F_path        string `default:"path"`
	F_ext         string `default:"ext"`
	F_type        string `default:"type"`
}

// 初始化
func (f *FileModel) Init() *FileModel {
	// functions.ReflectModel(&f.Field)
	// f.Model = base.NewModel(f.Field.T_table)
	return f
}

// 新增数据
func (f *FileModel) Insert(data map[string]interface{}) int {
	// ......
	// result,_ := f.Model.Data(data).Insert()

	result := 2
	return result
}

// 查重
func (f *FileModel) FindRepeatFile(md5 string, sha256 string, size int64) map[string]interface{} {
	rs := make(map[string]interface{})
	// .....
	// rs := f.Model.Where(map[string]interface{}{
	// 	f.Field.F_md5: md5,
	// 	f.Field.F_sha256: sha256,
	// 	f.Field.F_size: size,
	// }).Find()

	return rs
}
