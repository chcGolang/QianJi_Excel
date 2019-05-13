package service

import (
	"QianJi_Excel/dao"
	"QianJi_Excel/model"
	"QianJi_Excel/utils"
	"QianJi_Excel/vo"
)

type IConsumptionTypeService interface {
	// 查询全部的消费类型
	FindConsumptionTypeAll()(voInfo []vo.ConsumptionTypeVo,err error)
	// 分页查询
	FindConsumptionTypeToPage(consumption_name string,page,pageSize int)(voList []vo.ConsumptionTypeInfoVo,count int,err error)
	// 根据id查询
	FindConsumptionTypeById(id int)(voInfo vo.ConsumptionTypeVo,err error)
	// 根据Id更新
	UpdateConsumptionTypeById(voInfo vo.ConsumptionTypeVo)(err error)
	// 添加消费类型
	SaveConsumptionType(voInfo vo.ConsumptionTypeSaveVo)(err error)
	// 根据id删除消费类型,同时删除关联的规则
	DeleteConsumptionTypeById(id int)(err error)
}

func NewConsumptionTypeService()IConsumptionTypeService  {
	return &consumptionTypeService{
		ConsumptionTypeDao:dao.NewConsumptionTypeDao(),
	}
}

type consumptionTypeService struct {
	ConsumptionTypeDao dao.ConsumptionTypeDao
}

func (c *consumptionTypeService) DeleteConsumptionTypeById(id int) (err error) {
	return c.ConsumptionTypeDao.DeleteById(id)
}

func (c *consumptionTypeService) SaveConsumptionType(voInfo vo.ConsumptionTypeSaveVo) (err error) {
	// 校验类型是否存在
	if err = NewValidateService().CheckConsumptionTypeName(voInfo.ConsumptionName);err != nil{
		return
	}
	consumptionType := model.ConsumptionType{}
	if err = utils.DeepCopy(&consumptionType, voInfo);err != nil{
		return
	}
	return c.ConsumptionTypeDao.Save(&consumptionType)
}

func (c *consumptionTypeService) UpdateConsumptionTypeById(voInfo vo.ConsumptionTypeVo) (err error) {
	return c.ConsumptionTypeDao.UpdateById(voInfo.Id,voInfo.ConsumptionName)
}

func (c *consumptionTypeService) FindConsumptionTypeById(id int) (voInfo vo.ConsumptionTypeVo, err error) {
	return c.ConsumptionTypeDao.FindById(id)
}

func (c *consumptionTypeService) FindConsumptionTypeToPage(consumption_name string,page, pageSize int) (voList []vo.ConsumptionTypeInfoVo,count int, err error) {
	assemblePage(&page,&pageSize)
	return c.ConsumptionTypeDao.FindToPage(consumption_name,page,pageSize)
}

func (c *consumptionTypeService) FindConsumptionTypeAll() (voInfo []vo.ConsumptionTypeVo, err error) {
	return c.ConsumptionTypeDao.FindConsumptionTypeAll()
}

