package controller

import (
	"QianJi_Excel/constant"
	"QianJi_Excel/service"
	"QianJi_Excel/vo"
	"github.com/kataras/iris"
)

type MatchingRuleController struct {
	Ctx iris.Context
	MatchingRuleService service.IMatchingRuleService
	ConsumptionTypeService service.IConsumptionTypeService
	RulesTypeService service.IRulesTypeService
	ValidateService service.IValidateService
}


// 查询规则信息列表
func (m *MatchingRuleController)PostListdata()  resultResp {

	csv_type:= m.Ctx.PostValueIntDefault("csv_type",0)
	page:=m.Ctx.PostValueIntDefault("page",1)
	pageSize:=m.Ctx.PostValueIntDefault("pageSize",4)
	rulesTypeId:=m.Ctx.PostValueIntDefault("rulesTypeId",-1)
	consumptionTypeId:=m.Ctx.PostValueIntDefault("consumptionTypeId",-1)
	valueData := m.Ctx.PostValueDefault("valueData","")
	ckeckPage(&page,&pageSize)
	resultList,count, err := m.MatchingRuleService.FindTypeMatchingRuleBypage(csv_type,rulesTypeId,consumptionTypeId,valueData,page,pageSize)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesDataByPage(resultList,page,pageSize,count)

}

// 查询csv文件类型
func (m *MatchingRuleController)GetCsv_type_info() resultResp  {
	csvTypeInfo := make([]vo.CsvTypeInfoVo,2)
	csvTypeInfo[0] = vo.CsvTypeInfoVo{
		CsvType:constant.ALIPAY_CSV_TYPE,
		CsvName:"支付宝Csv",
	}
	csvTypeInfo[1] = vo.CsvTypeInfoVo{
		CsvType:constant.WECHAT_CSV_TYPE,
		CsvName:"微信Csv",
	}
	return succeesData(csvTypeInfo)
}

// 查询csv标题
func (m *MatchingRuleController)GetRules_info()resultResp  {
	csvType, _ := m.Ctx.URLParamInt("csv_type")
	voInfo, err := m.RulesTypeService.FindByCsvType(csvType)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesData(voInfo)
}

// 获取钱迹消费类型
func (m *MatchingRuleController)GetConsumption_type() resultResp {
	voInfo, err := m.ConsumptionTypeService.FindConsumptionTypeAll()
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesData(voInfo)
}

// 查询匹配值查询方式
func (m *MatchingRuleController)GetFuzzy()resultResp{
	voList := make([]vo.FuzzyInfoVo,2)
	voList[0] = vo.FuzzyInfoVo{
		Fuzzyid:constant.FUZZY_FALSE,
		FuzzyName:constant.FuzzyNameMap[constant.FUZZY_FALSE],
	}
	voList[1] = vo.FuzzyInfoVo{
		Fuzzyid:constant.FUZZY_TRUE,
		FuzzyName:constant.FuzzyNameMap[constant.FUZZY_TRUE],
	}
	return succeesData(voList)
}

// 添加规则
func (m *MatchingRuleController)PostSave_matching_rule()resultResp{
	var saveVo = vo.TypeMatchingRuleSaveVo{}
	if err := m.Ctx.ReadForm(&saveVo);err != nil{
		return failError(err,fali_param)
	}
	if err := m.MatchingRuleService.SaveTypeMatchingRuleInfo(saveVo);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("添加成功")
}

// 根据Id查询规则详情
func (m *MatchingRuleController)GetFind_id()resultResp{
	id, err := m.Ctx.URLParamInt("id")
	if err != nil{
		return failError(err,fali_param)
	}
	result, err := m.MatchingRuleService.FindTypeMatchingRuleById(id)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesData(result)
}
// 更新规则
func (m *MatchingRuleController)PostUpdate_matching_rule()resultResp{
	editVo := vo.TypeMatchingRuleEditVo{}
	if err := m.Ctx.ReadJSON(&editVo);err != nil{
		return failError(err,fali_param)
	}
	// 参数校验
	if errList := m.ValidateService.Validators(editVo);len(errList)>0{
		return failError(errList[0],fali_param)
	}
	if err := m.MatchingRuleService.UpdateTypeMatchingRuleById(editVo);err != nil{
		return failError(err,fail_service)
	}
	return succeesData("修改成功")
}

// 根据Id删除规则
func (m *MatchingRuleController)PostDelete_id()resultResp  {
	id, err := m.Ctx.PostValueInt("id")
	if err != nil{
		return failError(err,fali_param)
	}
	err = m.MatchingRuleService.DeleteTypeMatchingRuleById(id)
	if err != nil{
		return failError(err,fail_service)
	}
	return succeesData("删除成功")
}
