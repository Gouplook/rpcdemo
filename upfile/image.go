/**
 * @Author: yinjinlin
 * @File:  image
 * @Description:
 * @Date: 2021/5/25 上午10:22
 */

package upfile

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"github.com/h2non/filetype"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"rpcdemo/common/models"
	_const "rpcdemo/lang/const"
	"rpcdemo/rpcinterface/interface/file"
	"rpcdemo/tools"
	"rpcdemo/upbase/common/toolLib"
	"strconv"
	"strings"
	"time"
)

type ImageLogic struct {
}

// 获取图片的hash值
func (i *ImageLogic) GetImageByHashs(hashs []string) (*[]map[string]interface{}, error) {
	// 从数据库中查找，返回

	return nil, nil
}

// 根据URL地址获取远程图片
func (i *ImageLogic) SaveImgFromUrl(urlStr string) (reply file.ReplyFileInfo, err error) {

	reply = file.ReplyFileInfo{
		Id:   0,
		Hash: "",
		Path: "",
	}

	// 检查URL是否正确
	isUrl := false
	if strings.Contains(urlStr, "http://") {
		isUrl = true
	} else if strings.Contains(urlStr, "https://") {
		isUrl = true
	}
	if !isUrl {
		return
	}



	// request参数
	request := http.Request{}
	parse, _ := url.Parse(urlStr)
	request.URL = parse
	request.Method = http.MethodGet

	// 模糊....
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// 链接clien
	client := http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}

	// 请求
	resp, err := client.Do(&request)
	defer resp.Body.Close()
	if resp.StatusCode != 200 || err != nil {
		return
	}

	body := resp.Body
	imgBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	// 对数据进行加密
	objMd5 := md5.New()
	objMd5.Write(imgBytes)
	strMd5 := hex.EncodeToString(objMd5.Sum(nil))

	objSha256 := sha256.New()
	objSha256.Write(imgBytes)
	strSha256 := hex.EncodeToString(objSha256.Sum(nil))
	imgSize, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)

	// 保存之前查找一下，是否存在。
	imageModel := new(models.ImageModel).Init()
	res := imageModel.FindRepeatImage(strMd5,strSha256,imgSize)
	if len(res) != 0 {
		reply.Id, _ = strconv.Atoi(res[imageModel.Field.F_id].(string))
		reply.Hash = res[imageModel.Field.F_hash].(string)
		reply.Path = "配置文件路径" + reply.Hash
		return
	}

	// go get github.com/satori/go.uuid
	u2 := uuid.NewV4()

	// 如何计算hash值
	hash := strings.Replace(u2.String(), "-", "", -1)
	fileName := tools.GetFileNameByHash(hash)
	path := filepath.Join("../xxx", tools.GetFileNameByHash(hash))

	// creates a directory named path
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {

	}
	// 打开文件
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return
	}

	// 文件写入
	if _, err = f.Write(imgBytes); err != nil {
		return
	}
	fileNames := strings.Split("/", urlStr)
	fileExt := resp.Header.Get("Content-Type")
	if len(fileExt) == 0 {
		fileExt = "png"
	}
	// fileName 为何-1
	name := tools.GetFileName(fileNames[len(fileNames)-1])

	// 加密后的数据插入数据库

	id := imageModel.Insert(map[string]interface{}{
		imageModel.Field.F_name:        name,
		imageModel.Field.F_md5:         strMd5,
		imageModel.Field.F_sha256:      strSha256,
		imageModel.Field.F_size:        imgSize,
		imageModel.Field.F_type:        1, // 上传图片类型，自定义
		imageModel.Field.F_ext:         "png",
		imageModel.Field.F_path:        fileName,
		imageModel.Field.F_hash:        hash,
		imageModel.Field.F_create_time: time.Now().Unix(),
	})
	if id != 0 {
		reply.Id = id
		reply.Hash = hash
		reply.Path = "配置文件路径" + reply.Hash

		return
	}

	return
}

// 上传图片
func (i *ImageLogic)UploadImage(image *file.ArgsFile, reply *file.ReplyFileInfo) (err error) {

	// 判断是否是图片
	if !filetype.IsImage(image.Context) {
		// 不是图片
		return  toolLib.CreateKcErr(_const.IS_NOT_IMAGE)
	}

	// 数据加密
	objMd5 := md5.New()
	objMd5.Write(image.Context)
	strMd5 := hex.EncodeToString(objMd5.Sum(nil))

	objSha256 := sha256.New()
	objSha256.Write(image.Context)
	strSha256 := hex.EncodeToString(objSha256.Sum(nil))
	imageModel := new(models.ImageModel).Init()
	res := imageModel.FindRepeatImage(strMd5, strSha256, image.FileHeader.Size)

	// 数据库里，就不用上传了
	if len(res) != 0{
		//
		reply.Id, _ = strconv.Atoi(res[imageModel.Field.F_id].(string))
		reply.Hash = res[imageModel.Field.F_hash].(string)
		// reply.Path = kcgin.AppConfig.String("upload.image.host") + reply.Hash
		reply.Path = "配置文件路径" + reply.Hash
		return nil
	}
	//推断给定缓冲区的文件类型，检查其幻数签名
	kind,_ := filetype.Match(image.Context)
	if kind == filetype.Unknown {
		return  toolLib.CreateKcErr(_const.IS_NOT_IMAGE)
	}
	u2 := uuid.NewV4()
	hash := strings.Replace(u2.String(),"-","",-1)
	filename := tools.GetFileNameByHash(hash)
	path := filepath.Join("../xxx", tools.GetFileNameByHash(hash))
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	// 打开文件
	f, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0644)
	defer  f.Close()
	if err != nil {
		return
	}
	if _, err = f.Write(image.Context);err != nil {
		return err

	}

	name := tools.GetFileName(image.FileHeader.Filename)

	id := imageModel.Insert(map[string]interface{}{
		imageModel.Field.F_name:        name,
		imageModel.Field.F_size:        image.FileHeader.Size,
		imageModel.Field.F_sha256:      strSha256,
		imageModel.Field.F_md5:         strMd5,
		imageModel.Field.F_type:        image.Type,
		imageModel.Field.F_ext:         kind.Extension,
		imageModel.Field.F_path:        filename,
		imageModel.Field.F_hash:        hash,
		imageModel.Field.F_create_time: time.Now().Unix(),

	})
	if id != 0{
		reply.Id = id
		reply.Hash = hash
		reply.Path = "../" + reply.Hash
		return nil
	}

	return toolLib.CreateKcErr(_const.UPLOAD_FAIL)
}