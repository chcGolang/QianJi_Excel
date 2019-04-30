package service

import "QianJi_Excel/constant"

// 组装分页
func assemblePage(page,pageSize *int)  {
	*page = (*page-1)* (*pageSize)
}

// 校验csvType参数
func checkCsvType(csvType int)int  {
	if csvType!=constant.WECHAT_CSV_TYPE && csvType!=constant.ALIPAY_CSV_TYPE{
		csvType = 0
	}
	return csvType
}
