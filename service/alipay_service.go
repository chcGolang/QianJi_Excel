package service


import (
"QianJi_Excel/utils"
"bytes"
"encoding/json"
"io/ioutil"
)

type AlipayData struct {
	TransactionCreateTime string `json:"交易创建时间"`
	TransactionNum string `json:"交易号"`
	TransactionCounterparty string `json:"交易对方"`
	TransactionSourcePlace string `json:"交易来源地"`
	TransactionStatus string `json:"交易状态"`
	PayTime string `json:"付款时间"`
	CommodityName string `json:"商品名称"`
	OrderNum string `json:"商家订单号"`
	Remarks string `json:"备注"`
	Refund string `json:"成功退款(元)"`
	BalanceOrpayments string `json:"收/支"`
	LastUpdateTime string `json:"最近修改时间"`
	CoverCharge string `json:"服务费(元)"`
	Type string `json:"类型"`
	FundsStatus string `json:"资金状态"`
	Money string `json:"金额(元)"`
}

// 解析支付宝excel的数据
func ReadAlipayExcel(excelPath string) (alipayDataList []AlipayData,err error) {
	byteList, err := ioutil.ReadFile(excelPath)
	if err != nil {
		return
	}
	byteList, err = utils.Decode(byteList)
	if err != nil {
		return
	}
	reader := bytes.NewReader(byteList)
	// 解析csv文件数据
	csvMap := utils.ReadCsvToMap(reader, 16)
	// 转成AlipayData数据
	marshal, err := json.Marshal(csvMap)
	if err != nil {
		return
	}
	if err = json.NewDecoder(bytes.NewReader(marshal)).Decode(&alipayDataList);err != nil{
		return
	}

	return
}


