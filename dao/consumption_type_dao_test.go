package dao

import (
	"QianJi_Excel/model"
	"testing"
)

func TestConsumptionTypeDaoImpl_Save(t *testing.T) {
	consumptionTypeDaoI := NewConsumptionTypeDao()
	consumptionType := model.ConsumptionType{
		ConsumptionName: "出行",
	}
	if err :=consumptionTypeDaoI.Save(&consumptionType);err != nil{
		panic(err)
	}
}
