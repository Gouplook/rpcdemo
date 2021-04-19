项目开发总结


1： 获取详情 --- select
getInfo := getBusIs(busId)  // 获取的
// 临时存放的
baseInfo := struct { 
BusId       int
Sales       int
.....
}{}
// 转换时注意，
mapstructure.WeakDecode(getInfo, &baseInfo)