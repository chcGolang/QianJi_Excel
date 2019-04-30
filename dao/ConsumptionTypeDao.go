package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"QianJi_Excel/vo"
	"github.com/jinzhu/gorm"
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
}

func NewConsumptionTypeDao()(consumptionTypeDaoI ConsumptionTypeDao)  {
	return &consumptionTypeDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}

type consumptionTypeDaoImpl struct {
	db *gorm.DB
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
