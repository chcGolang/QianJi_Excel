package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"QianJi_Excel/vo"
	"github.com/jinzhu/gorm"
	"time"
)

type ConsumptionTypeDao interface {
	// 添加消费类型
	Save(consumptionType *model.ConsumptionType)error
	// 根据Id批量删除
	DeleteByIds(ids []int)error
	// 查询全部的消费类型
	FindConsumptionTypeAll()(voInfo []vo.ConsumptionTypeVo,err error)
	// 查询消费类型id是否存在
	CheckConsumptionTypeId(id int)(check bool,err error)
	// 分页查询
	FindToPage(consumption_name string,page,pageSize int)(voList []vo.ConsumptionTypeInfoVo,count int,err error)
	// 根据Id查询
	FindById(id int)(voInfo vo.ConsumptionTypeVo,err error)
	// 根据Id删除,并且删除关联的规则
	DeleteById(id int)(err error)
	// 根据Id更新
	UpdateById(id int,consumption_name string)(err error)
	// 根据消费类型名称查询
	FindConsumptionNameToCount(consumption_name string)(count int,err error)
}

func NewConsumptionTypeDao()(consumptionTypeDaoI ConsumptionTypeDao)  {
	return &consumptionTypeDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}

type consumptionTypeDaoImpl struct {
	db *gorm.DB
}

func (c *consumptionTypeDaoImpl) FindConsumptionNameToCount(consumption_name string) (count int, err error) {
	err = c.db.Table(model.ConsumptionType{}.TableName()).
		Where("consumption_name = ?", consumption_name).Count(&count).Error
	return
}

func (c *consumptionTypeDaoImpl) UpdateById(id int, consumption_name string)(err error){
	mapUpdate := make(map[string]interface{})
	mapUpdate["consumption_name"] = consumption_name
	mapUpdate["updated_at"] = time.Now()
	return c.db.Table(model.ConsumptionType{}.TableName()).Where("id = ?",id).Updates(mapUpdate).Error
}

func (c *consumptionTypeDaoImpl) DeleteById(id int) (err error) {
	tx := c.db.Begin()
	defer db.RecoverRollback(&err,tx)
	if errList := tx.Delete(model.ConsumptionType{}, "id = ?", id).GetErrors();len(errList) > 0{
		tx.Rollback()
		return errList[0]
	}
	if err = tx.Delete(model.TypeMatchingRule{},"consumption_type_id = ?",id).Error;err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

func (c *consumptionTypeDaoImpl) FindById(id int) (voInfo vo.ConsumptionTypeVo, err error) {
	err = c.db.Table(model.ConsumptionType{}.TableName()).
		Where("id = ?", id).Find(&voInfo).Error
	return
}

func (c *consumptionTypeDaoImpl) FindToPage(consumption_name string,page, pageSize int) (voList []vo.ConsumptionTypeInfoVo, count int,err error) {
	table := c.db.Table(model.ConsumptionType{}.TableName())
	if len(consumption_name)>0{
		table = table.Where("consumption_name LIKE ?","%"+consumption_name+"%")
	}
	err = table.Order("updated_at DESC").Count(&count).
		Limit(pageSize).Offset(page).Find(&voList).Error
	return
}

func (c *consumptionTypeDaoImpl) CheckConsumptionTypeId(id int) (check bool, err error) {
	var count int
	err = c.db.Table(model.ConsumptionType{}.TableName()).Where("id = ?",id).Count(&count).Error
	check = true
	if count == 0{
		check = false
	}
	return
}

func (c *consumptionTypeDaoImpl) FindConsumptionTypeAll() (voInfo []vo.ConsumptionTypeVo, err error) {
	err = c.db.Table(model.ConsumptionType{}.TableName()).Find(&voInfo).Error
	return
}

func (c *consumptionTypeDaoImpl) DeleteByIds(ids []int) error {
	return c.db.Delete(model.ConsumptionType{},"id IN (?)",ids).Error
}

func (c *consumptionTypeDaoImpl) Save(consumptionType *model.ConsumptionType) error {
	return c.db.Save(consumptionType).Error
}
