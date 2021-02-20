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

type ArgsFile struct {
	Type       int    //  上传文件类型
	Context    []byte // 文件内容
	FileHeader *multipart.FileHeader
}

type ReplyFileInfo struct {
	Id   int
	Hash string
	Path string
}

// 上传照片接口
type Userinfo interface {
	// 图片上传

	//根据url地址获取远程图片
	SaveImgFromUrl(ctx context.Context, imgStr *string, reply *ReplyFileInfo ) error
}
