/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/20 11:05
@Description:

*********************************************/
package file

import (
	"context"
	"mime/multipart"
)


// ===============文件/图片上传===========
// 上传入参
type ArgsFile struct {
	Type       int    //  上传文件类型
	Context    []byte // 文件内容
	FileHeader *multipart.FileHeader
}

// 返回文件信息
type ReplyFileInfo struct {
	Id   int
	Hash string
	Path string
}

// 文件/图片上传接口
type Userinfo interface {
	// 文件上传
	UploadFile(ctx context.Context,args *ArgsFile,reply *ReplyFileInfo) error

	// 图片上传
	UploadImage(ctx context.Context, image *ArgsFile, reply *ReplyFileInfo) error
	//根据URL地址获取远程图片
	SaveImgFromUrl(ctx context.Context, imgStr *string, reply *ReplyFileInfo ) error
}
