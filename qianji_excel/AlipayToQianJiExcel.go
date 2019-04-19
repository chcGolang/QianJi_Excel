package qianji_excel

import (
	"QianJi_Excel/alipay_excel"
	"QianJi_Excel/constant"
	"QianJi_Excel/utils"
)

const(
	Counterparty_Lease = "信用借还服务商"
	Counterparty_elm = "饿了么"
	Counterparty_didi = "滴滴出行"
	Counterparty_YuE = "天弘基金管理有限公司"
	SourcePlace_TaoBao = "淘宝"
	CommodityName_Prepare = "代付"

)

const(
	Effective_TransactionStatus = "交易成功"
	Counterparty_HuaHe = "花呗"
	Balance_Status = "收入"
	Payments_Status = "支出"
)

// 将支付宝账单转账钱迹的excel数据
func AlipayDataToExcelMap(aliPayDataList []alipay_excel.AlipayData) (dataExcelMap []map[string]interface{},err error) {
	var(
		aliPayData alipay_excel.AlipayData
		dataMap map[string]interface{}
	)
	dataExcelMap = make([]map[string]interface{},0)
	rules, err := typeMatchingRuleDaoI.FindTypeMatchingRuleInfoByCsvType(constant.ALIPAY_CSV_TYPE)
	if err != nil{
		return
	}
	for _,aliPayData = range aliPayDataList{

		if aliPayData.TransactionStatus != Effective_TransactionStatus{
			continue
		}

		if aliPayData.TransactionCounterparty == Counterparty_HuaHe {
			continue
		}

		if aliPayData.BalanceOrpayments != Balance_Status &&
			aliPayData.BalanceOrpayments != Payments_Status{
			continue
		}
		dataMap = make(map[string]interface{})
		dataMapTag, err := utils.GetMapTagToStruct(aliPayData)
		if err != nil{
			return
		}
		dataMap[QianJiExcel_Type] = typeToQianjiType(dataMapTag,rules)
		dataMap[QianJiExcel_BalanceOfPayments] = aliPayData.BalanceOrpayments
		dataMap[QianJiExcel_Money] = aliPayData.Money
		dataMap[QianJiExcel_Time] = aliPayData.TransactionCreateTime
		dataMap[QianJiExcel_Remarks] = aliPayData.CommodityName
		dataExcelMap = append(dataExcelMap,dataMap)
	}
	return
}

