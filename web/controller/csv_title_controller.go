package controller

import (
	"QianJi_Excel/service"
	"QianJi_Excel/vo"
	"github.com/kataras/iris"
)

// csv标题管理
type CsvTitleController struct {
	Ctx iris.Context
	RulesTypeService service.IRulesTypeService
	ValidateService service.IValidateService
}

// 查询csv标题列表
func (c *CsvTitleController)PostList()resultResp  {
	csv_type := c.Ctx.PostValueIntDefault("csv_type", 0)
	page := c.Ctx.PostValueIntDefault("page", 1)
	pageSize := c.Ctx.PostValueIntDefault("pageSize", 10)
	titleName := c.Ctx.PostValueDefault("titleName", "")
	voInfo,count, err := c.RulesTypeService.FindRulesTypeToPage(titleName, csv_type, page, pageSize)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesDataByPage(voInfo,page,pageSize,count)
}

func (c *CsvTitleController)GetByid()resultResp  {
	id, err := c.Ctx.URLParamInt("id")
	if err != nil{
		return failError(err,fali_param)
	}
	voInfo, err := c.RulesTypeService.FindRulesTypeInfoById(id)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesData(voInfo)
}

func (c *CsvTitleController)PostUpdate_byid()resultResp  {
	id, err := c.Ctx.PostValueInt("id")
	if err != nil{
		return failError(err,fali_param)
	}
	rules_name := c.Ctx.PostValue("rules_name")
	csv_type, err := c.Ctx.PostValueInt("csv_type")
	if err != nil{
		return failError(err,fali_param)
	}
	if err = c.RulesTypeService.UpdateById(id,csv_type,rules_name);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("操作成功")
}

func (c *CsvTitleController)PostAdd_rulestype()resultResp  {
	saveVo := vo.RulesTypeSaveVo{}
	if err := c.Ctx.ReadJSON(&saveVo);err != nil{
		return failError(err,fali_param)
	}
	// 校验参数
	if errList := c.ValidateService.Validators(saveVo);len(errList)>0{
		return failError(errList[0],fali_param)
	}

	if err := c.RulesTypeService.SaveRulesType(saveVo);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("操作成功")
}

func (c *CsvTitleController)PostDelete_id()resultResp  {
	id, err := c.Ctx.PostValueInt("id")
	if err != nil{
		return failError(err,fali_param)
	}
	if err = c.RulesTypeService.DeleteRulesTypeById(id);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("")
}
