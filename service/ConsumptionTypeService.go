package service

import (
	"QianJi_Excel/dao"
	"QianJi_Excel/vo"
)

type IConsumptionTypeService interface {
	// 查询全部的消费类型
	FindConsumptionTypeAll()(voInfo []vo.ConsumptionTypeVo,err error)
}

func NewConsumptionTypeService()IConsumptionTypeService  {
	return &consumptionTypeService{
		ConsumptionTypeDao:dao.NewConsumptionTypeDao(),
	}
}

type consumptionTypeService struct {
	ConsumptionTypeDao dao.ConsumptionTypeDao
}

func (c *consumptionTypeService) FindConsumptionTypeAll() (voInfo []vo.ConsumptionTypeVo, err error) {
	return c.ConsumptionTypeDao.FindConsumptionTypeAll()
}

