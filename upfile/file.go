/**
 * @Author: yinjinlin
 * @File:  file
 * @Description:
 * @Date: 2021/5/25 上午10:22
 */

package upfile

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/h2non/filetype"
	uuid "github.com/satori/go.uuid"
	"os"
	"path/filepath"
	"rpcdemo/common/models"
	_const "rpcdemo/lang/const"
	"rpcdemo/rpcinterface/interface/file"
	"rpcdemo/tools"
	"rpcdemo/upbase/common/functions"
	"rpcdemo/upbase/common/toolLib"
	"rpcdemo/upgin/logs"
	"strings"
	"time"
)

type FileLogic struct {
}


func (f *FileLogic)UploadFile(argFile *file.ArgsFile, reply *file.ReplyFileInfo) error {

	// 缓存中的文件类型
	kind, _ := filetype.Match(argFile.Context)
	if kind == filetype.Unknown {
		// 未知文件类型
		return toolLib.CreateKcErr(_const.FILE_TYPE_ERR)
	}
	// 附件文件类型
	accExt := "./xx"   // 本地配置文件中定义
	accExts := strings.Split(accExt, ",")

	// 检索附件扩展类型
	if !functions.InArray(kind.Extension, accExts) {
		logs.Info(2)
		return toolLib.CreateKcErr(_const.FILE_TYPE_ERR)
	}

	// 数据加密
	objMd5 := md5.New()
	objMd5.Write(argFile.Context)
	strMd5 := hex.EncodeToString(objMd5.Sum(nil))

	objSha256 := sha256.New()
	objSha256.Write(argFile.Context)
	strSha256 := hex.EncodeToString(objSha256.Sum(nil))

	fileModel := new(models.FileModel).Init()
	res := fileModel.FindRepeatFile(strMd5, strSha256, argFile.FileHeader.Size)
	if len(res) != 0 {
		reply.Hash = res[fileModel.Field.F_hash].(string)
		reply.Path = "./xx 配置文件" + reply.Hash + "." + res[fileModel.Field.F_ext].(string)
		return nil
	}

	// or error handling
	u2 := uuid.NewV4()
	hash := strings.ReplaceAll(u2.String(), "-", "")
	filename := tools.GetFileNameByHash(hash) + "." + kind.Extension
	path := filepath.Join("./xx 配置文件", filename)
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	osFile,err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0644)

	defer osFile.Close()
	if err != nil {
		return err
	}
	if _, err = osFile.Write(argFile.Context); err != nil {
		return err
	}

	name := tools.GetFileName(argFile.FileHeader.Filename)
	id := fileModel.Insert(map[string]interface{}{
		fileModel.Field.F_name:        name,
		fileModel.Field.F_size:        argFile.FileHeader.Size,
		fileModel.Field.F_sha256:      strSha256,
		fileModel.Field.F_md5:         strMd5,
		fileModel.Field.F_type:        argFile.Type,
		fileModel.Field.F_ext:         kind.Extension,
		fileModel.Field.F_path:        filename,
		fileModel.Field.F_hash:        hash,
		fileModel.Field.F_create_time: time.Now().Unix(),
	})
	if id != 0 {
		reply.Hash = hash
		reply.Path = "./xx 配置文件" + reply.Hash + "." + kind.Extension
		return nil
	}
	return toolLib.CreateKcErr(_const.UPLOAD_FAIL)





}

