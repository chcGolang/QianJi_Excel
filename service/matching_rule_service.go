package service

import (
	"QianJi_Excel/dao"
	"QianJi_Excel/model"
	"QianJi_Excel/utils"
	"QianJi_Excel/vo"
	"strings"
)

// 规则操作服务
type IMatchingRuleService interface {
	// 查询匹配规则列表
	FindTypeMatchingRuleBypage(csvType,rulesTypeId,consumptionTypeId int, valueData string, page,pageSize int)(resultList []vo.TypeMatchingRuleVo,count int,err error)
	// 添加规则
	SaveTypeMatchingRuleInfo(saveVo vo.TypeMatchingRuleSaveVo)error
	// 根据Id查询规则信息
	FindTypeMatchingRuleById(id int)(result vo.TypeMatchingRuleInfoVo,err error)
	// 根据id更新规则
	UpdateTypeMatchingRuleById(editVo vo.TypeMatchingRuleEditVo)(err error)
	//根据Id删除规则
	DeleteTypeMatchingRuleById(id int)(err error)
}
func NewMatchingRuleService() IMatchingRuleService {
	return &matchingRuleService{
		typeMatchingRuleDao:dao.NewTypeMatchingRuleDao(),
	}
}
type matchingRuleService struct {
	typeMatchingRuleDao dao.TypeMatchingRuleDao
}

func (v *matchingRuleService) DeleteTypeMatchingRuleById(id int) (err error) {
	return v.typeMatchingRuleDao.DeleteById(id)
}

func (v *matchingRuleService) UpdateTypeMatchingRuleById(editVo vo.TypeMatchingRuleEditVo) (err error) {
	err = NewValidateService().CheckRulesTypeIdByCsvType(editVo.CsvType, editVo.RulesTypeId)
	if err != nil{
		return
	}
	var typeMatchingRule model.TypeMatchingRule
	if err = utils.DeepCopy(&typeMatchingRule, editVo);err != nil{
		return
	}
	return v.typeMatchingRuleDao.UpdateTypeMatchingRuleById(typeMatchingRule)

}

func (v *matchingRuleService) FindTypeMatchingRuleById(id int) (result vo.TypeMatchingRuleInfoVo, err error) {
	return v.typeMatchingRuleDao.FindById(id)
}

func (v *matchingRuleService) SaveTypeMatchingRuleInfo(saveVo vo.TypeMatchingRuleSaveVo) error {
	var matchingRule = model.TypeMatchingRule{}
	// 校验文件类型
	if checkCsvType(saveVo.CsvType) == 0{
		return utils.Errorf("csv文件类型异常:csvType=%d",saveVo.CsvType)
	}
	// 校验文件类型与csv表标题是否匹配
	if err := NewValidateService().
		CheckRulesTypeIdByCsvType(saveVo.CsvType, saveVo.RulesTypeId);err != nil{
			return err
	}
	// 校验消费类型是否正确
	if err := NewValidateService().CheckConsumptionTypeId(saveVo.ConsumptionTypeId);err != nil{
		return err
	}

	if err := utils.DeepCopy(&matchingRule, saveVo);err != nil{
		return err
	}
	return v.typeMatchingRuleDao.Save(&matchingRule)
}

func (v *matchingRuleService) FindTypeMatchingRuleBypage(csvType,rulesTypeId,consumptionTypeId int, valueData string, page,pageSize int)(resultList []vo.TypeMatchingRuleVo,count int,err error) {
	assemblePage(&page,&pageSize)
	csvType = checkCsvType(csvType)
	valueData = strings.Trim(valueData," ")
	return v.typeMatchingRuleDao.FindTypeMatchingRule(csvType,rulesTypeId,consumptionTypeId,valueData,page,pageSize)
}
