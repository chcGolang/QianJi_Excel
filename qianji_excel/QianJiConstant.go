package qianji_excel

import (
	"QianJi_Excel/constant"
	"QianJi_Excel/dao"
	"QianJi_Excel/vo"
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
