/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 09:49
@Description: 区域接口

*********************************************/
package public

import "context"

//区域返回
type RegionInfo struct {
	RegionId int
	RegionName string
}
// 获取城市/区街道入参
type ArgsRegion struct {
	RegionId int
}

type Region interface {
	// 获取省份/直辖市
	GetProvince(ctx context.Context, reply *[]RegionInfo) error
	// 获取城市/区街道
	GetRegion(ctx context.Context, args *ArgsRegion, reply *[]RegionInfo) error
	// 获取单条区域信息
	GetByRegionId(ctx context.Context,args *ArgsRegion,reply *RegionInfo) error

}
