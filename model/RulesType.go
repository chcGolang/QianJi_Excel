package model

// csv字段表
type RulesType struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// csv标题名称
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
	// 1:支付宝,2:微信
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
}

func (RulesType) TableName() string {
	return "rules_type"
}
