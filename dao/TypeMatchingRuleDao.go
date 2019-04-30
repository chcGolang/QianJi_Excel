package dao

import (
	"QianJi_Excel/constant"
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"QianJi_Excel/vo"
	"fmt"
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
	FindTypeMatchingRule(csvType,rulesTypeId,consumptionTypeId int, valueData string, page,pageSize int)(resultList []vo.TypeMatchingRuleVo,count int,err error)
	// 根据Id查询
	FindById(id int)(result vo.TypeMatchingRuleInfoVo,err error)
	// 根据id更新
	UpdateTypeMatchingRuleById(editVo model.TypeMatchingRule)(err error)
	// 根据Id删除
	DeleteById(id int)(err error)
}

func NewTypeMatchingRuleDao() (typeMatchingRuleDaoI TypeMatchingRuleDao) {
	return &typeMatchingRuleDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}

type typeMatchingRuleDaoImpl struct {
	db *gorm.DB
}

func (t *typeMatchingRuleDaoImpl) DeleteById(id int) (err error) {
	return t.db.Delete(model.TypeMatchingRule{}, "id = ?", id).Error
}

func (t *typeMatchingRuleDaoImpl) UpdateTypeMatchingRuleById(editVo model.TypeMatchingRule) (err error) {
	return t.db.Save(&editVo).Error
}

func (t *typeMatchingRuleDaoImpl) FindById(id int) (result vo.TypeMatchingRuleInfoVo, err error) {
	err = t.db.Table("type_matching_rule").
		Where("id = ?",id).
		Find(&result).Error
	return
}

func (t *typeMatchingRuleDaoImpl) FindTypeMatchingRule(csvType,rulesTypeId,consumptionTypeId int, valueData string, page,pageSize int) (resultList []vo.TypeMatchingRuleVo,count int, err error) {
	resultList = make([]vo.TypeMatchingRuleVo,0)
	sqlSelect := fmt.Sprintf("tmr.id,tmr.csv_type,tmr.value_data,rt.rules_name,ct.consumption_name," +
		"CASE tmr.fuzzy WHEN %d then '%s' WHEN %d THEN '%s' ELSE '未知' END AS fuzzy_name",constant.FUZZY_FALSE,
		constant.FuzzyNameMap[constant.FUZZY_FALSE],constant.FUZZY_TRUE,constant.FuzzyNameMap[constant.FUZZY_TRUE])
	db := t.db.Table("type_matching_rule tmr").
		Select(sqlSelect).
		Joins("INNER JOIN rules_type rt ON rt.id = tmr.rules_type_id").
		Joins("INNER JOIN consumption_type ct ON ct.id = tmr.consumption_type_id")
	sqlWhere := "0 = ?"
	sqlValue := make([]interface{},0)
	sqlValue = append(sqlValue,0)
	if csvType!=0{
		sqlWhere+=" AND tmr.csv_type = ?"
		sqlValue = append(sqlValue,csvType)
	}


	if rulesTypeId >= 0{
		sqlWhere+=fmt.Sprint(" AND tmr.rules_type_id = ?")
		sqlValue = append(sqlValue,rulesTypeId)
	}
	if consumptionTypeId >= 0{
		sqlWhere+=fmt.Sprint(" AND tmr.consumption_type_id = ?")
		sqlValue = append(sqlValue,consumptionTypeId)
	}
	if len(valueData) > 0{

		sqlWhere+=fmt.Sprint(" AND tmr.value_data LIKE ?")
		sqlValue = append(sqlValue,"%"+valueData+"%")

	}
	err = db.Where(sqlWhere,sqlValue...).Count(&count).
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
