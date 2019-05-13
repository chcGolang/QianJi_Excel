package model

import "time"

type CsvFileInfo struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// 上传文件存放地址
	UploadFilePath string `json:"upload_file_path" form:"upload_file_path" query:"upload_file_path"`
	// 转换成钱迹模板的文件存放地址
	QianjiFilePath string `json:"qianji_file_path" form:"qianji_file_path" query:"qianji_file_path"`
	// 账单所属时间
	BillTime *time.Time `json:"bill_time" form:"bill_time" query:"bill_time"`
	// 状态(0:上传成功,1:转换信息录入成功,2:转换成功,-1:上传失败,-2:转换信息录入失败.-3:转换失败)
	Status int `json:"status" form:"status" query:"status"`
	// 失败原因
	ErrMsg string `json:"err_msg" form:"err_msg" query:"err_msg"`
	// 1:支付宝,2:微信
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// 添加时间
	CreatedAt time.Time `json:"created_at" form:"created_at" query:"created_at"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" query:"updated_at"`



}

func (CsvFileInfo)TableName()string  {
	return "csv_file_info"
}
