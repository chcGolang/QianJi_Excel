package constant

const(
	ALIPAY_CSV_TYPE = 1
	WECHAT_CSV_TYPE = 2
)

var CsvTypeName =map[int]string{
	ALIPAY_CSV_TYPE:"支付宝csv文件",
	WECHAT_CSV_TYPE:"微信csv文件",
}
