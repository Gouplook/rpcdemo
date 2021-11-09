
>>>-----log-1
//////////////////////////////--->>>>rpcInsurance>>--///////////////////////////////////
开发日期：2020-12-2 

开发需求 ：
1: 获取保单任务信息
2: 获取续保任务信息


添加/修改源码： 
1： rpcInsurance/logics/insruanceLogic.go
2： common/logics/insruanceLogic.go:184 

添加/修改文档：
无


>>>-----log-2
////////////////////////-----rpcVisualization---/////////////////////////////////////////
开发日期：2021-2-2 

开发需求：
1:  更改商户风险等级和安全码信息（2021-02-02）
2： 添加监管可视化 保险保单信息 （2020-12-02）

添加/修改源码：
1 ： logic/BusRiskLogic.go
2 ： common/logics/cardPackagePolicyLogic.go:26


添加/修改交换机或队列源码：
1：consumeTask:  dataVisualization/busRiskRandSafeCode.go
2：consumeTask:  insurance/cardPackPolicyTask.go:43

添加/修改文档：
无



>>>-----log-3
///////////////////////////---rpcComtreeData----//////////////////////////////////////

开发日期：2021-3-30

开发需求：
1:  添加 预付卡消费 信息
2:  添加 预付卡保险出单 信息
3:  添加 预付卡交易 信息
4:  添加 全国市场规模 信息
5:  添加 业务规模 信息


添加/修改源码： 
1： common/logics/consumptionLogic.go:30
2： common/logics/consumptionLogic.go:104
3： common/logics/consumptionLogic.go:158
4： common/logics/consumptionLogic.go:256
5： common/logics/consumptionLogic.go:444

添加/修改交换机或队列源码：
1： order/consumeSucNumAndCardAmountCount.go:43
2： insurance/cardPackPolicyTask.go:51
3： order/statisticsOrderPaySuc.go:41
4： busAddToEs.go:54
5： busAddToEs.go:58

添加/修改文档：
无



>>>-----log-4
///////////////////////--rpcStaff---/////////////////////

开发日期：2020-10-16

开发需求：
1: 添加招聘信息（系列）

添加/修改源码：
1： common/logics/StaffRecruitLogic.go:25

添加/修改文档：
http://apidoc.shutung.com/web/#/4?page_id=494



>>>-----log-5
////////////////////////--rpcComment---//////////////////////////

开发日期：2020-09-25

开发需求：
1: 用户回答提问 (2020-09-25)
2: 获取客户评价数据（客户今日完成评价总数目， 待处理总数目）(2020-10-23)

添加/修改源码：
1: common/logics/AskAnswerLogic.go:27
2: common/logics/DayStatisticsLogic.go:19

添加/修改文档：
无

>>>-----log-6
////////////////////////--rpcFinancial---//////////////////////////

开发日期：2020-12-25

开发需求：
1: 保险公司上月保费金额 （2020-12-25）
2: 保险公司累计保费金额

---Api接口 系列 ----
3：获取保险公司商户统计(当前银行， 所有银行商户总数）（2020-12-28）
4：获取保险公司月保费金额（2020-12-28）

添加/修改源码：
1: common/logics/InsuranceLogic.go:38
2: common/logics/InsuranceLogic.go:98

添加/修改交换机或队列源码：
1：insurance/insuranceAssetsMonth.go:38
2：insurance/insuranceAssetsTotal.go:38


添加/修改文档：
无

>>>-----log-7
> ////////////////////////--rpcCards---//////////////////////////
开发日期：

开发需求：
1: 购买或者充值须100起 (2020-11-20)
2: 获取所有卡的发布数量 rpc内部调用 （2020-11-16）
3：卡项收藏状态（2020-09-29）
4：添加充值卡（2020-11-6）
5：店铺新增充值规则  充值卡规则编辑  获取充值卡规则详情....（2020-10-29）

添加/修改源码：
1: common/logics/CardLogic.go:282  
2: common/logics/CardLogic.go:1540
3：common/logics/ItemLogic.go:2801
4：common/logics/RcardLogic.go:88
5：common/logics/RcardLogic.go:1350

添加/修改文档：
无


>>>-----log-8
> ////////////////////////--rpcorder---//////////////////////////
开发日期：

开发需求：
1: 增加待处理 待提货单数量 (2020-10-10)
2: 获取卡项目购卡总数（对应的店铺) (2020-11-13)
3: 统计用户复购率 (2020-11-18)

添加/修改源码：
1: common/logics/DayStatisticsLogic.go:993
2: common/logics/ItemOrderLogic.go:592
3: common/logics/OrderPayLogic.go:2478


添加/修改文档：
无




>>>-----log-9
> ////////////////////////--预约服务rpcReservation---//////////////////////////
开发日期：

开发需求：
1: 

添加/修改源码：
1: 

添加/修改文档：
无

>>>-----log-10
> ////////////////////////--rpcCards---//////////////////////////
开发日期：2021-04-12


开发需求：
1:单项目 综合卡 套餐  限时卡  限次卡  限时限次卡 充值卡 身份卡 一期优化
发布，上下架，设置适用门店

添加/修改源码：
1:common/logics/CardLogic.go
2：common/logics/HcardLogic.go
3：common/logics/HNCardLogic.go
....系列

添加/修改文档：
1：卡项 商品服务接口



sql:备注

kc_card:综合卡
ALTER TABLE `jkd_cards`.`kc_card`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;
ADD COLUMN `is_permanent_validity` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否永久有效：1-是；2-否；如果永久有效则service_period字段可忽略，否者必填' AFTER `service_period`;

kc_shop_card:已添加到门店的综合卡
ALTER TABLE `jkd_cards`.`kc_shop_card`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_hcard:限时卡表
ALTER TABLE `jkd_cards`.`kc_hcard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_shop_hcard:已添加到门店的限时卡
ALTER TABLE `jkd_cards`.`kc_shop_hcard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_hncard:限时限次卡表
ALTER TABLE `jkd_cards`.`kc_hncard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_shop_hncard: 已添加到门店的限时限次卡
ALTER TABLE `jkd_cards`.`kc_shop_hncard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_icard:身份卡表
ALTER TABLE `jkd_cards`.`kc_icard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_shop_icard:已添加到门店的身份卡
ALTER TABLE `jkd_cards`.`kc_shop_icard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_hncard:限次卡表
ALTER TABLE `jkd_cards`.`kc_hncard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;
ADD COLUMN `is_permanent_validity` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否永久有效：1-是；2-否；如果永久有效则service_period字段可忽略，否者必填' AFTER `service_period`;

kc_shop_ncard:已添加到门店的限次卡
ALTER TABLE `jkd_cards`.`kc_shop_ncard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;


kc_rcard：充值卡
ALTER TABLE `jkd_cards`.`kc_rcard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;
ADD COLUMN `is_permanent_validity` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否永久有效：1-是；2-否；如果永久有效则service_period字段可忽略，否者必填' AFTER `service_period`;
ADD COLUMN `is_have_discount` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否享受折扣：1-无折扣；2-有折扣',



kc_shop_rcard：添加到门店的充值卡
ALTER TABLE `jkd_cards`.`kc_shop_rcard`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_sm: 套餐表
ALTER TABLE `jkd_cards`.`kc_sm`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;
ADD COLUMN `is_permanent_validity` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否永久有效：1-是；2-否；如果永久有效则service_period字段可忽略，否者必填' AFTER `service_period`;

kc_shop_sm:已添加到门店的套餐
ALTER TABLE `jkd_cards`.`kc_shop_sm`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_single:单项目基础信息表
ALTER TABLE `jkd_cards`.`kc_single`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;

kc_shop_single: 已添加到门店的单项目
ALTER TABLE `jkd_cards`.`kc_shop_single`
ADD COLUMN `is_del` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0=否，1=是' AFTER `ctime`,
ADD COLUMN `del_time` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间' AFTER `is_del`;



>>>-----log-11
> ////////////////////////--rpcRisk---////////////////////////// 
开发日期：2021-04-12


开发需求：
1：添加 已兑付金额
2: 添加 预付卡风险管理系统--会员总人数
3: 添加 预付卡风险管理系统--风控系统商户ID 门店总数量统计
4: 添加 预付卡风险管理系统--会员复购率
5: 添加 预付卡风险管理系统 (表 ：BusBasicAccumulative)
.....

添加/修改源码：
1：common/logics/BusBasicAccumulativeLogic.go:65


添加/修改文档：
1：无文档修改


>>>-----log-12
> ////////////////////////--rpcxxx---//////////////////////////
开发日期：2021-04-12


开发需求：
1：

添加/修改源码：
1：



添加/修改文档：
1：