package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"github.com/jinzhu/gorm"
)

const RulesTypeName  = "rules_type"
type RulesTypeDao interface {
	// 添加
	Save(rulesType *model.RulesType)error
	// 根据Id批量删除
	DeleteByIds(ids []int)error
}
func NewRulesTypeDao() (rulesTypeDaoI RulesTypeDao){
	return &rulesTypeDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}
type rulesTypeDaoImpl struct {
	db *gorm.DB
}

func (r *rulesTypeDaoImpl) DeleteByIds(ids []int) error {
	return r.db.Delete(model.RulesType{}, "id IN (?)", ids).Error

}

func (r *rulesTypeDaoImpl) Save(rulesType *model.RulesType) error {
	save := r.db.Table(RulesTypeName).Save(rulesType)
	return save.Error
}


