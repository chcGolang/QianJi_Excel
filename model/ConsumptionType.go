package model

// 消费类型
type ConsumptionType struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// 类型名称
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name"`
}

func (ConsumptionType)TableName()string  {
	return "consumption_type"
}
