package dao

import (
	"QianJi_Excel/constant"
	"QianJi_Excel/model"
	"fmt"
	"testing"
)

func TestTypeMatchingRuleDaoImpl_Save(t *testing.T) {
	typeMatchingRuleDaoI := NewTypeMatchingRuleDao()
	rule := model.TypeMatchingRule{
		CsvType:constant.ALIPAY_CSV_TYPE,
		RulesTypeId:4,
		Fuzzy:0,
		ValueData:"淘宝",
		ConsumptionTypeId:10,
	}
	if err := typeMatchingRuleDaoI.Save(&rule);err != nil{
		panic(err)
	}

}

func TestTypeMatchingRuleDaoImpl_FindTypeMatchingRuleInfoByCsvType(t *testing.T) {
	typeMatchingRuleDaoI := NewTypeMatchingRuleDao()
	rules, err := typeMatchingRuleDaoI.FindTypeMatchingRuleInfoByCsvType(constant.ALIPAY_CSV_TYPE)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v \n",rules)
}
