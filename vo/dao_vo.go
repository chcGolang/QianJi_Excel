package vo

type TypeMatchingRule struct {
	// csv类型
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// 是否模糊匹配(0:不模糊,1:模糊)
	Fuzzy int `json:"fuzzy" form:"fuzzy" query:"fuzzy"`
	// 匹配值
	Value string `json:"value" form:"value" query:"value"`
	// 类型名称
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name"`
	// csv标题名称
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
}
