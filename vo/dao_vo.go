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

type RulesTypeInfoVo struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// csv标题名称
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
}

type ConsumptionTypeVo struct {
	Id int `json:"id" form:"id" query:"id"`
	// 类型名称
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name" validate:"required"`
}

type ConsumptionTypeSaveVo struct {
	// 类型名称
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name" validate:"required"`
}

type ConsumptionTypeInfoVo struct {
	Id int `json:"id" form:"id" query:"id"`
	// 类型名称
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name"`
	// 添加时间
	CreatedAt string `json:"created_at" form:"created_at" query:"created_at"`
	// 修改时间
	UpdatedAt string `json:"updated_at" form:"updated_at" query:"updated_at"`
}


type RulesTypeInfoPageVo struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// csv标题名称
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
	// 1:支付宝,2:微信
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// csv类型名称
	CsvTypeName string `json:"csv_type_name" form:"csv_type_name" query:"csv_type_name"`

}

type RulesTypeFindVo struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// csv标题名称
	RulesName string `json:"rules_name" form:"rules_name" query:"rules_name"`
	// 1:支付宝,2:微信
	CsvType int `json:"csv_type" form:"csv_type" query:"csv_type"`
	// csv类型名称
	CsvTypeName string `json:"csv_type_name" form:"csv_type_name" query:"csv_type_name"`

}
