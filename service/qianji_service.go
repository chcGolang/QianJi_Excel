package service

import (
	"QianJi_Excel/constant"
	"QianJi_Excel/dao"
	"QianJi_Excel/vo"
	"QianJi_Excel/wechatpay_excel"
	"fmt"
	"strings"
)
const(
	type_Lease = "信用借还"
	type_TaoBaoPay = "淘宝购物"
	type_ToGo = "外卖"
	type_Traffic = "出行"
	type_Financial = "理财"
	type_Other = "其他"
	type_Prepare = "代付"
	type_Transfer = "转账"
	type_Res = "红包"
	type_Entertainment = "娱乐"
)

const(
	QianJiExcel_Time = "时间"
	QianJiExcel_Type = "分类"
	QianJiExcel_BalanceOfPayments = "收入或支出"
	QianJiExcel_Money = "金额"
	QianJiExcel_Remarks = "备注"
)

var ExcelQianJiKey = []string{QianJiExcel_Time,QianJiExcel_Type,QianJiExcel_BalanceOfPayments,QianJiExcel_Money,QianJiExcel_Remarks}


var typeMatchingRuleDaoI = dao.NewTypeMatchingRuleDao()

// 转换支付类型
func typeToQianjiType(dataMap map[string]string,typeMatchingRuleList []vo.TypeMatchingRule)(typeStr string)  {
	for mapKey,mapValue:= range dataMap{
		for _,ruleValue:=range typeMatchingRuleList{
			if mapKey == ruleValue.RulesName{
				switch ruleValue.Fuzzy {
				case constant.FUZZY_FALSE:
					if ruleValue.Value == mapValue{
						typeStr = ruleValue.ConsumptionName
						return
					}
				case constant.FUZZY_TRUE:
					if strings.Contains(mapValue,ruleValue.Value){
						typeStr = ruleValue.ConsumptionName
						return
					}
				}
			}
		}
	}
	typeStr =  type_Other
	return
}

const(
	Wechat_Balance_Status = "收入"
	Wechat_Payments_Status = "支出"
)

const(
	Wechat_Type_Transfer = "转账"
	Wechat_Type_Red = "微信红包"
	Wechat_Type_Cash = "零钱提现"
	Wechat_Counterparty_didi = "滴滴出行"
	Wechat_Counterparty_Entertainment = "充值服务"
)

// 将微信账单转账钱迹的excel数据
func WechatPayDataToExcelMap(wechatPayDataList []wechatpay_excel.WechatPayData) []map[string]interface{} {
	var(
		wechatPayData wechatpay_excel.WechatPayData
		dataMap map[string]interface{}
		dataExcelMap = make([]map[string]interface{},0)
	)
	for _,wechatPayData = range wechatPayDataList{


		if wechatPayData.BalanceOrpayments != Wechat_Balance_Status &&
			wechatPayData.BalanceOrpayments != Wechat_Payments_Status{
			continue
		}

		if wechatPayData.Type == Wechat_Type_Cash{
			continue
		}
		dataMap = make(map[string]interface{})
		dataMap[QianJiExcel_Type] = wechatTypeToQianjiType(wechatPayData)
		dataMap[QianJiExcel_BalanceOfPayments] = wechatPayData.BalanceOrpayments
		dataMap[QianJiExcel_Money] = strings.Trim(wechatPayData.Money,"¥")
		dataMap[QianJiExcel_Time] = wechatPayData.TransactionCreateTime
		dataMap[QianJiExcel_Remarks] = fmt.Sprintf("%s-%s",wechatPayData.TransactionCounterparty,wechatPayData.CommodityName)
		dataExcelMap = append(dataExcelMap,dataMap)
	}
	return dataExcelMap
}

// 转换支付类型
func wechatTypeToQianjiType(wechatPayData wechatpay_excel.WechatPayData)(typeStr string)  {
	if wechatPayData.Type == Wechat_Type_Transfer{
		typeStr = type_Transfer
	}else if wechatPayData.Type == Wechat_Type_Red{
		typeStr = type_Res
	}else if wechatPayData.TransactionCounterparty == Wechat_Counterparty_didi{
		typeStr = type_Traffic
	}else if wechatPayData.TransactionCounterparty == Wechat_Counterparty_Entertainment{
		typeStr = type_Entertainment
	}else{
		typeStr = type_Other
	}

	return
}
