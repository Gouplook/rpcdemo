/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/4 10:36
@Description:

*********************************************/
package poplar

import "context"

type ArgsGetStudentByName struct {
	Name string // 姓名，必填
}
type ArgsGetStudentsByAge struct {
	Age int // 年龄，必填
	//common.Paging
}
type ArgsGetStudents struct {
	//common.Paging
}
type ReplyStudent struct {
	Name string // 姓名
	Age int     // 年龄
}
type Student interface {
	// 根据姓名获取单条数据
	GetStudentByName(ctx context.Context, args *ArgsGetStudentByName, reply *ReplyStudent) error
	// 根据年龄获取数据
	GetStudentsByAge(ctx context.Context, args *ArgsGetStudentsByAge, reply *[]ReplyStudent) error
	// 获取所有数据
	GetStudents(ctx context.Context, args *ArgsGetStudents, reply *[]ReplyStudent) error
}