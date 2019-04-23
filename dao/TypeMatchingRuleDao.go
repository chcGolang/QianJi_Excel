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
	// 查询匹配规则
	FindTypeMatchingRule(csvTypr int,page,pageSize int)(resultList []vo.TypeMatchingRuleVo,count int,err error)
}

func NewTypeMatchingRuleDao() (typeMatchingRuleDaoI TypeMatchingRuleDao) {
	return &typeMatchingRuleDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}

type typeMatchingRuleDaoImpl struct {
	db *gorm.DB
}

func (t *typeMatchingRuleDaoImpl) FindTypeMatchingRule(csvTypr int,page,pageSize int) (resultList []vo.TypeMatchingRuleVo,count int, err error) {
	resultList = make([]vo.TypeMatchingRuleVo,0)
	err = t.db.Table("type_matching_rule tmr").
		Select("tmr.id,tmr.csv_type,tmr.value_data,rt.rules_name,ct.consumption_name").
		Where("tmr.csv_type = ?", csvTypr).
		Joins("INNER JOIN rules_type rt ON rt.id = tmr.rules_type_id").
		Joins("INNER JOIN consumption_type ct ON ct.id = tmr.consumption_type_id").
		Count(&count).
		Limit(pageSize).Offset(page).
		Find(&resultList).Error
	return
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
