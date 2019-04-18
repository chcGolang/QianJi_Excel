package qianji_excel

import (
	"QianJi_Excel/wechatpay_excel"
	"fmt"
	"strings"
)

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
