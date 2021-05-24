/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/19 16:44
@Description:

*********************************************/
package models

import (
	"rpcdemo/upbase/common/functions"
	"rpcdemo/upbase/common/models/base"
	"rpcdemo/upgin/orm"
)

//表结构体
type DemoModel struct {
	Model   *base.Model
	Field DemoModelField
}

//表字段
type DemoModelField struct {
	T_table string
	F_reminder string `default:"reminder"`


	F_CreatedAt string `default:"CreatedAt"`
	CreatedAt string

}

func (d *DemoModel)Init(ormer ...orm.Ormer)*DemoModel{
	functions.ReflectModel(&d.Field)
	d.Model = base.NewModel(d.Field.T_table, ormer...)

	return d
}

// 新增数据
func (m *DemoModel) Insert(data map[string]interface{}) int {
	result, _ := m.Model.Data(data).Insert()
	return result
}

//批量添加
func (m *DemoModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := m.Model.InsertAll(data)
	return result
}

// 更新数据
func (m *DemoModel) Update(where, data map[string]interface{}) bool {
	if len(where) == 0 {
		return false
	}
	_, err := m.Model.Where(where).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// 单条数据查询
func (m *DemoModel) Find(where map[string]interface{}) map[string]interface{} {
	if len(where) == 0 {
		return make(map[string]interface{})
	}
	return m.Model.Where(where).Find()
}

// 基础查询（多条）
func (m *DemoModel) Select(where map[string]interface{}) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	return m.Model.Where(where).Select()
}

// 带分页查询 （多条）
func (m *DemoModel) SelectByPage(where map[string]interface{}, start, limit int) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}

	// 需要修改，刷选结果CreatedAt作为排序的字段
	return m.Model.Where(where).Limit(start, limit).OrderBy(m.Field.CreatedAt + " DESC ").Select()
}




