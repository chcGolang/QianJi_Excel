package qianji_excel

import (
	"QianJi_Excel/alipay_excel"
	"strings"
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
func AlipayDataToExcelMap(aliPayDataList []alipay_excel.AlipayData) []map[string]interface{} {
	var(
		aliPayData alipay_excel.AlipayData
		dataMap map[string]interface{}
		dataExcelMap = make([]map[string]interface{},0)
	)
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
		dataMap[QianJiExcel_Type] = alipayTypeToQianjiType(aliPayData)
		dataMap[QianJiExcel_BalanceOfPayments] = aliPayData.BalanceOrpayments
		dataMap[QianJiExcel_Money] = aliPayData.Money
		dataMap[QianJiExcel_Time] = aliPayData.TransactionCreateTime
		dataMap[QianJiExcel_Remarks] = aliPayData.CommodityName
		dataExcelMap = append(dataExcelMap,dataMap)
	}
	return dataExcelMap
}

// 转换支付类型
func alipayTypeToQianjiType(alipayData alipay_excel.AlipayData)(typeStr string)  {
	if strings.Contains(alipayData.TransactionCounterparty,Counterparty_Lease){
		typeStr = type_Lease
	}else if alipayData.TransactionCounterparty == Counterparty_elm {
		typeStr = type_ToGo
	}else if alipayData.TransactionCounterparty == Counterparty_didi {
		typeStr = type_Traffic
	}else if alipayData.TransactionCounterparty == Counterparty_YuE {
		typeStr = type_Financial
	}else if alipayData.CommodityName == CommodityName_Prepare {
		typeStr = type_Prepare
	}else if alipayData.TransactionSourcePlace == SourcePlace_TaoBao{
		typeStr = type_TaoBaoPay
	}else{
		typeStr = type_Other
	}

	return
}
