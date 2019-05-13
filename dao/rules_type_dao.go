package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"QianJi_Excel/vo"
	"github.com/jinzhu/gorm"
	"time"
)

const RulesTypeName  = "rules_type"
type RulesTypeDao interface {
	// 添加
	Save(rulesType *model.RulesType) error
	// 根据Id批量删除
	DeleteByIds(ids []int) error
	// 根据csv文件类型查询标题
	FindByCsvType(csvType int) (voInfo []vo.RulesTypeInfoVo, err error)
	// 根据csv文件类型查询标题是否存在
	FindRulesCountByCsvType(csvType int, rulesTypeId int) (checkBol bool, err error)
	// 分页查询csv标题列表
	FindToPage(titleName string, csv_type int, page, pageSize int) (voList []vo.RulesTypeInfoPageVo,count int, err error)
	// 根据Id查询
	FindById(id int)(voInfo vo.RulesTypeFindVo,err error)
	// 根据Id更新
	UpdateById(id int,rules_name string)(err error)
	// 根据csv文件类型和标题名称查询条目数
	FindCountByCsvTypeAndRulesName(rules_name string,csv_type int)(count int,err error)
	// 根据Id删除,并且删除TypeMatchingRule关联的数据
	DeleteById(id int)(err error)
}
func NewRulesTypeDao() (rulesTypeDaoI RulesTypeDao){
	return &rulesTypeDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}
type rulesTypeDaoImpl struct {
	db *gorm.DB
}

func (r *rulesTypeDaoImpl) DeleteById(id int) (err error) {
	tx := r.db.Begin()
	defer db.RecoverRollback(&err,tx)
	if err = tx.Delete(model.RulesType{},"id = ?", id).Error;err != nil{
		tx.Rollback()
		return
	}

	if err = tx.Delete(model.TypeMatchingRule{},"rules_type_id = ?", id).Error;err != nil{
		tx.Rollback()
		return
	}

	return tx.Commit().Error


}

func (r *rulesTypeDaoImpl) FindCountByCsvTypeAndRulesName(rules_name string, csv_type int) (count int, err error) {
	err = r.db.Table(model.RulesType{}.TableName()).Where("rules_name = ? AND csv_type = ?", rules_name, csv_type).
		Count(&count).Error
	return
}

func (r *rulesTypeDaoImpl) UpdateById(id int, rules_name string)(err error) {
	updateMap := make(map[string]interface{})
	updateMap["rules_name"]=rules_name
	updateMap["updated_at"] = time.Now()
	return r.db.Table(model.RulesType{}.TableName()).
		Where("id = ?",id).
		Updates(updateMap).Error
}

func (r *rulesTypeDaoImpl) FindById(id int) (voInfo vo.RulesTypeFindVo, err error) {
	err = r.db.Table(model.RulesType{}.TableName()).Where("id = ?",id).Find(&voInfo).Error
	return
}

func (r *rulesTypeDaoImpl) FindToPage(titleName string, csv_type int, page, pageSize int) (voList []vo.RulesTypeInfoPageVo,count int, err error) {
	sqlWhere := "0=?"
	sqlValue := make([]interface{},0)
	sqlValue=append(sqlValue,0)
	if len(titleName)>0{
		sqlWhere += " AND rules_name LIKE ?"
		sqlValue=append(sqlValue,"%"+titleName+"%")
	}
	if csv_type != 0{
		sqlWhere += " AND csv_type = ?"
		sqlValue=append(sqlValue,csv_type)
	}
	err = r.db.Table(model.RulesType{}.TableName()).Where(sqlWhere,sqlValue...).
		Order("updated_at DESC" +
			",id DESC").
		Count(&count).
		Limit(pageSize).
		Offset(page).
		Find(&voList).Error
	return
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


