/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/19 17:29
@Description:

*********************************************/
package base

import (
	"rpcdemo/upgin"
	"rpcdemo/upgin/logs"
	"rpcdemo/upgin/orm"
	"strconv"
	"fmt"
)

var(
	err error
)

//初始化驱动
func init(){
	logs.Info("Init driver.go mysql start")
	//设置驱动数据库连接参数
	dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",upgin.AppConfig.String("db.user"),upgin.AppConfig.String("db.pwd"),
		upgin.AppConfig.String("db.host"),upgin.AppConfig.String("db.port"),upgin.AppConfig.String("db.name"),upgin.AppConfig.String("db.charset"))
	//打印连接数据库参数
	logs.Info("DatabaseDriverConnect String:",dataSource)
	maxIdle,_:= strconv.Atoi(upgin.AppConfig.DefaultString("db.maxidle","10"))
	maxConn,_:= strconv.Atoi(upgin.AppConfig.DefaultString("db.maxconn","0"))
	maxTime := upgin.AppConfig.DefaultInt("db.maxlifetime", 10800)
	logs.Info("connMaxLifeTime(s)：", maxTime)
	//设置注册数据库
	if err == nil{
		err = orm.RegisterDataBase("default", upgin.AppConfig.String("db.type"), dataSource,maxIdle,maxConn,maxTime)
	}
}