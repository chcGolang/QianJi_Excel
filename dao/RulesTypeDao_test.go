package dao

import (
	"QianJi_Excel/model"
	"testing"
)

func TestRulesTypeDaoImpl_Save(t *testing.T) {

	//var rulesTypeDao RulesTypeDao
	rulesTypeDao := NewRulesTypeDao()
	rulesType := model.RulesType{
		RulesName:"出行",
		CsvType:1,
	}
	var err error
	if err = rulesTypeDao.Save(&rulesType);err!=nil{
		panic(err)
	}
}
