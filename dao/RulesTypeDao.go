package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"QianJi_Excel/vo"
	"github.com/jinzhu/gorm"
)

const RulesTypeName  = "rules_type"
type RulesTypeDao interface {
	// 添加
	Save(rulesType *model.RulesType)error
	// 根据Id批量删除
	DeleteByIds(ids []int)error
	// 根据csv文件类型查询标题
	FindByCsvType(csvType int )(voInfo []vo.RulesTypeInfoVo,err error)
	// 根据csv文件类型查询标题是否存在
	FindRulesCountByCsvType(csvType int,rulesTypeId int)(checkBol bool,err error)
}
func NewRulesTypeDao() (rulesTypeDaoI RulesTypeDao){
	return &rulesTypeDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}
type rulesTypeDaoImpl struct {
	db *gorm.DB
}

func (r *rulesTypeDaoImpl) FindRulesCountByCsvType(csvType int, rulesTypeId int) (checkBol bool, err error) {
	var count int
	err = r.db.Table(model.RulesType{}.TableName()).Where("csv_type = ? AND id = ?", csvType, rulesTypeId).Count(&count).Error
	checkBol = true
	if count == 0{
		checkBol = false
	}
	return

}

func (r *rulesTypeDaoImpl) FindByCsvType(csvType int) (voInfo []vo.RulesTypeInfoVo, err error) {
	table := r.db.Table(model.RulesType{}.TableName())
	if csvType!=0{
		table = table.Where("csv_type = ? ",csvType)
	}
	err = table.Find(&voInfo).Error
	return
}

func (r *rulesTypeDaoImpl) DeleteByIds(ids []int) error {
	return r.db.Delete(model.RulesType{}, "id IN (?)", ids).Error

}

func (r *rulesTypeDaoImpl) Save(rulesType *model.RulesType) error {
	save := r.db.Table(RulesTypeName).Save(rulesType)
	return save.Error
}


