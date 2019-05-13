package controller

import (
	"QianJi_Excel/service"
	"QianJi_Excel/vo"
	"github.com/kataras/iris"
)

type ConsumptionTypeController struct {
	Ctx iris.Context
	ConsumptionTypeService service.IConsumptionTypeService
	ValidateService service.IValidateService
}
// 分页查询
func (c *ConsumptionTypeController)PostFind_page()resultResp  {
	page := c.Ctx.PostValueIntDefault("page", 1)
	pageSize := c.Ctx.PostValueIntDefault("pageSize", 10)
	consumption_name := c.Ctx.PostValueDefault("consumption_name", "")
	voList, count, err := c.ConsumptionTypeService.FindConsumptionTypeToPage(consumption_name,page, pageSize)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesDataByPage(voList,page,pageSize,count)
}
// 根据Id查询
func (c *ConsumptionTypeController)GetByid()resultResp  {
	id, err := c.Ctx.URLParamInt("id")
	if err != nil{
		return failError(err,fali_param)
	}
	voInfo, err := c.ConsumptionTypeService.FindConsumptionTypeById(id)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesData(voInfo)
}

// 根据Id更新
func (c *ConsumptionTypeController)PostUpdate_byid()resultResp  {
	c.Ctx.Values()
	voInfo:= vo.ConsumptionTypeVo{}
	if err := c.Ctx.ReadJSON(&voInfo);err != nil{
		return failError(err,fali_param)
	}
	if errList := c.ValidateService.Validators(voInfo);len(errList)>0{
		return failError(errList[0],fali_param)
	}
	if err := c.ConsumptionTypeService.UpdateConsumptionTypeById(voInfo);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("")

}

// 添加消费类型
func (c *ConsumptionTypeController)PostSave()resultResp  {
	voInfo := vo.ConsumptionTypeSaveVo{}
	if err := c.Ctx.ReadJSON(&voInfo);err != nil{
		return failError(err,fali_param)
	}
	if errList := c.ValidateService.Validators(voInfo);len(errList)>0{
		return failError(errList[0],fali_param)
	}
	if err := c.ConsumptionTypeService.SaveConsumptionType(voInfo);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("")
}

// 根据id删除消费类型
func (c *ConsumptionTypeController)PostDelete_byid()resultResp  {
	id, err := c.Ctx.PostValueInt("id")
	if err != nil{
		return failError(err,fali_param)
	}
	if err = c.ConsumptionTypeService.DeleteConsumptionTypeById(id);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("")
}
