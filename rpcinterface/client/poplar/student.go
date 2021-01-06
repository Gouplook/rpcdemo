/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/1/4 11:02
@Description:

*********************************************/
package poplar

import (
	"context"
	"rpcdemo/rpcinterface/interface/poplar"
)

type Student struct {
	//client.Baseclient

}
func (s *Student)Init() *Student{
	// 配置服务Name和Path

	return s
}
func (s *Student)GetStudentByName(ctx context.Context, args *poplar.ArgsGetStudentByName, reply *poplar.ReplyStudent) error {
	//return student.Call(ctx, "GetStudentByName", args, reply)
	return nil
}
