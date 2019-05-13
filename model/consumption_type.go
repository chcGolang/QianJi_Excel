package model

import "time"

// 消费类型
type ConsumptionType struct {
	Id int `json:"id" form:"id" query:"id" gorm:"primary_key"`
	// 类型名称
	ConsumptionName string `json:"consumption_name" form:"consumption_name" query:"consumption_name"`
	// 添加时间
	CreatedAt time.Time `json:"created_at" form:"created_at" query:"created_at"`
	// 修改时间
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" query:"updated_at"`
}

func (ConsumptionType)TableName()string  {
	return "consumption_type"
}
