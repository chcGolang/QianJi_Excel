package vo

type TypeMatchingRuleVo struct {
	Id int `json:"id" form:"id" query:"id"`
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	ValueData string `json:"value_data" form:"value_data" query:"value_data"`
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name"`
	FuzzyName string `json:"fuzzy_name" form:"fuzzy_name" query:"fuzzy_name"`
}

// csv的文件类型
type CsvTypeInfoVo struct {
	// id
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// 名称
	CsvName string `json:"csv_name" form:"csv_name" query:"csv_name"`
}

type FuzzyInfoVo struct {
	Fuzzyid int `json:"fuzzyid" form:"fuzzyid" query:"fuzzyid"`
	FuzzyName string `json:"fuzzy_name" form:"fuzzy_name" query:"fuzzy_name"`
}

type TypeMatchingRuleSaveVo struct {
	// csv类型(1:支付宝,2:微信)
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// 匹配字段Id
	RulesTypeId int `json:"rules_type_id" form:"rules_type_id" query:"rules_type_id"`
	// 是否模糊匹配(0:全文匹配,1:模糊)
	Fuzzy int `json:"fuzzy" form:"fuzzy" query:"fuzzy"`
	// 匹配值
	ValueData string `json:"value_data" form:"value_data" query:"value_data"`
	// 消费类型Id
	ConsumptionTypeId int `json:"consumption_type_id" form:"consumption_type_id" query:"consumption_type_id"`
}

type TypeMatchingRuleInfoVo struct {
	// 主键id
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// csv类型(1:支付宝,2:微信)
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// 匹配字段Id
	RulesTypeId int `json:"rules_type_id" form:"rules_type_id" query:"rules_type_id"`
	// 是否模糊匹配(0:全文匹配,1:模糊)
	Fuzzy int `json:"fuzzy" form:"fuzzy" query:"fuzzy"`
	// 匹配值
	ValueData string `json:"value_data" form:"value_data" query:"value_data"`
	// 消费类型Id
	ConsumptionTypeId int `json:"consumption_type_id" form:"consumption_type_id" query:"consumption_type_id"`
}

type TypeMatchingRuleEditVo struct {
	Id int `json:"id" form:"id" query:"id" validate:"checkMatchingRuleId"`
	// csv类型(1:支付宝,2:微信)
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type" validate:"csvType"`
	// 匹配字段Id
	RulesTypeId int `json:"rules_type_id" form:"rules_type_id" query:"rules_type_id"`
	// 是否模糊匹配(0:全文匹配,1:模糊)
	Fuzzy int `json:"fuzzy" form:"fuzzy" query:"fuzzy" validate:"checkFuzzy"`
	// 匹配值
	ValueData string `json:"value_data" form:"value_data" query:"value_data" validate:"checkValueData"`
	// 消费类型Id
	ConsumptionTypeId int `json:"consumption_type_id" form:"consumption_type_id" query:"consumption_type_id" validate:"consumptionTypeId"`
}
