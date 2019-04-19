package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"QianJi_Excel/vo"
	"github.com/jinzhu/gorm"
)


type TypeMatchingRuleDao interface {
	// 添加匹配规则
	Save(rule *model.TypeMatchingRule)error
	// 根据Id列表批量删除
	DeleteByIds(ids []int)error
	// 根据csvType查询规则信息
	FindTypeMatchingRuleInfoByCsvType(csvType int)([]vo.TypeMatchingRule,error)
}

func NewTypeMatchingRuleDao() (typeMatchingRuleDaoI TypeMatchingRuleDao) {
	return &typeMatchingRuleDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}

type typeMatchingRuleDaoImpl struct {
	db *gorm.DB
}

func (t *typeMatchingRuleDaoImpl) FindTypeMatchingRuleInfoByCsvType(csvType int) ([]vo.TypeMatchingRule,error) {
	var typeMatchingRuleInfo = make([]vo.TypeMatchingRule,0)
	scan := t.db.Table(model.TypeMatchingRule{}.TableName()+" tme").
		Select("tme.csv_type,tme.fuzzy,tme.value,ct.consumption_name,rt.rules_name").
		Joins("INNER JOIN consumption_type ct ON ct.id = tme.consumption_type_id").
		Joins("INNER JOIN rules_type rt ON rt.id = tme.rules_type_id").Scan(&typeMatchingRuleInfo)
	if scan.Error!=nil{
		return typeMatchingRuleInfo,scan.Error
	}

	return typeMatchingRuleInfo,nil
}

func (t *typeMatchingRuleDaoImpl) DeleteByIds(ids []int) error {
	return t.db.Delete(model.TypeMatchingRule{},"id IN (?)",ids).Error
}

func (t *typeMatchingRuleDaoImpl) Save(rule *model.TypeMatchingRule) error {
	return t.db.Save(rule).Error
}
