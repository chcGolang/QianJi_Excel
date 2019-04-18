package model

type TypeMatchingRule struct {
	// 主键id
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// csv类型(1:支付宝,2:微信)
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// 匹配字段Id
	RulesId int `json:"rules_id" form:"rules_id" query:"rules_id"`
	// 是否模糊匹配(0:不模糊,1:模糊)
	Fuzzy int `json:"fuzzy" form:"fuzzy" query:"fuzzy"`
	// 匹配值
	Value string `json:"value" form:"value" query:"value"`
	// 消费类型Id
	ConsumptionTypeId string `json:"consumption_type_id" form:"consumption_type_id" query:"consumption_type_id"`





}