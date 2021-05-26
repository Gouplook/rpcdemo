/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 10:09
@Description:

*********************************************/
package upregion

import (
	"context"
	"rpcdemo/rpcinterface/interface/public"
)

type RegionLogic struct {

}

// 获取单条区域信息
func (r *RegionLogic)GetByRegionId(ctx context.Context, args *public.ArgsRegion, reply *public.RegionInfo) error {
	// 1: 初始化mode
	// 2：根据区域ID（RegionId）在数据中查找
	// 3: 查找到数据进行返回

	return nil
}



// 20000+20000 = 40000 - 2700-1220-350-6000 =28000-10000 = 18000


//
