package wechatpay_excel

import (
	"QianJi_Excel/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type WechatPayData struct {
	TransactionNum string `json:"交易单号"`
	TransactionCounterparty string `json:"交易对方"`
	TransactionCreateTime string `json:"交易时间"`
	Type string `json:"交易类型"`
	CommodityName string `json:"商品"`
	OrderNum string `json:"商户单号"`
	Remarks string `json:"备注"`
	TransactionStatus string `json:"当前状态"`
	PaymentMode string `json:"支付方式"`
	BalanceOrpayments string `json:"收/支"`
	Money string `json:"金额(元)"`
}

func ReadWechatExcel(excelPath string)(wechatPayDataList []WechatPayData,err error)  {
	byteList, err := ioutil.ReadFile(excelPath)
	if err != nil {
		return
	}
	reader := bytes.NewReader(byteList)
	// 解析csv文件数据
	csvToMap := utils.ReadCsvToMap(reader, 11)
	marshal, err := json.Marshal(csvToMap)
	if err != nil {
		return
	}
	if err = json.NewDecoder(bytes.NewReader(marshal)).Decode(&wechatPayDataList);err != nil{
		return
	}
	return
}

