package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"github.com/jinzhu/gorm"
)

type ConsumptionTypeDao interface {
	// 添加消费类型
	Save(consumptionType *model.ConsumptionType)error
	// 根据Id批量删除
	DeleteByIds(ids []int)error
}

func NewConsumptionTypeDao()(consumptionTypeDaoI ConsumptionTypeDao)  {
	return &consumptionTypeDaoImpl{
		db:db.NewDataSource().MysqlDB,
	}
}

type consumptionTypeDaoImpl struct {
	db *gorm.DB
}

func (c *consumptionTypeDaoImpl) DeleteByIds(ids []int) error {
	return c.db.Delete(model.ConsumptionType{},"id IN (?)",ids).Error
}

func (c *consumptionTypeDaoImpl) Save(consumptionType *model.ConsumptionType) error {
	return c.db.Save(consumptionType).Error
}
