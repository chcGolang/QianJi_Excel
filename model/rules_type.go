package model

import "time"

// csv字段表
type RulesType struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// csv标题名称
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
	// 1:支付宝,2:微信
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// 添加时间
	CreatedAt time.Time `json:"created_at" form:"created_at" query:"created_at"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" query:"updated_at"`
}

func (RulesType) TableName() string {
	return "rules_type"
}
