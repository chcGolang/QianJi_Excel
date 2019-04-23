package vo

type TypeMatchingRuleVo struct {
	Id int `json:"id" form:"id" query:"id"`
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	ValueData string `json:"value_data" form:"value_data" query:"value_data"`
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name"`
}
