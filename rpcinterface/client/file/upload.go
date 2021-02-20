/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/20 13:09
@Description:

*********************************************/
package file

import "rpcdemo/rpcinterface/client"

type Upload struct {
	client.Baseclient
}

func (upload *Upload) Init() *Upload {
	upload.ServiceName = "rpc_file"
	upload.ServicePath = "Upload"
	return upload
}
