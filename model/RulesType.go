package model

type RulesType struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
}
