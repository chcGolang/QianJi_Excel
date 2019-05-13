package service

import (
	"QianJi_Excel/constant"
	"QianJi_Excel/dao"
	"QianJi_Excel/utils"
	"QianJi_Excel/vo"
	"github.com/smokezl/govalidators"
	"reflect"
	"strings"
)

// 参数校验服务
type IValidateService interface {
	// 根据csv文件类型校验csv表标题
	CheckRulesTypeIdByCsvType(csvType int,rulesTypeId int)(err error)
	// 校验钱迹消费类型
	CheckConsumptionTypeId(consumptionTypeId int)(err error)
	// 交运参数
	Validators(structWeb interface{})(errList []error)
	// 查询csv类型
	FindCsvTypeInfo()(voList []vo.CsvTypeInfoVo)
	// 校验RulesName是否存在
	CheckRulesName(rules_name string,csv_type int)(err error)
	// 校验钱迹消费类型名称
	CheckConsumptionTypeName(name string)(err error)
}

var validateServiceImpl = &validateService{
	RulesTypeDao: dao.NewRulesTypeDao(),
	ConsumptionTypeDao:dao.NewConsumptionTypeDao(),
	TypeMatchingRuleDao: dao.NewTypeMatchingRuleDao(),
}
func NewValidateService()IValidateService  {
	return validateServiceImpl
}
/*func newValidator() *govalidators.Validator {
	var validator = govalidators.New()
}*/

type validateService struct {
	dao.RulesTypeDao
	dao.ConsumptionTypeDao
	validator *govalidators.Validator
	dao.TypeMatchingRuleDao
}

func (v *validateService) CheckConsumptionTypeName(name string) (err error) {
	count := 0
	if count, err = v.ConsumptionTypeDao.FindConsumptionNameToCount(name);err != nil{
		return
	}
	if count >0 {
		err = utils.Errorf("消费类型已存在:%s",name)
	}
	return

}

func (v *validateService) CheckRulesName(rules_name string, csv_type int) (err error) {
	count, err := v.RulesTypeDao.FindCountByCsvTypeAndRulesName(rules_name, csv_type)
	if err !=nil{
		return
	}
	if count > 0 {
		return utils.Errorf("Csv标题已存在")
	}
	return nil
}

func (v *validateService) FindCsvTypeInfo() (voList []vo.CsvTypeInfoVo) {
	voList = make([]vo.CsvTypeInfoVo,2)
	voList[0] = vo.CsvTypeInfoVo{
		CsvType:constant.ALIPAY_CSV_TYPE,
		CsvName:"支付宝Csv",
	}
	voList[1] = vo.CsvTypeInfoVo{
		CsvType:constant.WECHAT_CSV_TYPE,
		CsvName:"微信Csv",
	}
	return
}

var validator = govalidators.New()
func (v *validateService) Validators(structObj interface{}) (errList []error) {
	validator.SetValidators(map[string]interface{}{
		"consumptionTypeId":v.validationConsumptionTypeId,
		"csvType":v.validationCsvType,
		"checkMatchingRuleId":v.validationMatchingRuleId,
		"checkFuzzy":v.validationFuzzy,
		"checkValueData":v.validationValueData,

	})
	errList = validator.Validate(structObj)

	return errList
}

// 消费类型校验
func (v *validateService)validationConsumptionTypeId(params map[string]interface{}, val reflect.Value, args ...string) (bool, error)  {
	if val.Kind() != reflect.Int{
		return false,utils.Errorf("钱迹消费类型的参数类型错误:%s",val.Kind().String())
	}
	id := val.Interface().(int)
	if err := v.CheckConsumptionTypeId(id);err != nil{
		return false,err
	}
	return true,nil

}

// 消费类型校验
func (v *validateService)validationValueData(params map[string]interface{}, val reflect.Value, args ...string) (bool, error)  {
	if val.Kind() != reflect.String{
		return false,utils.Errorf("匹配值的参数类型错误:%s",val.Kind().String())
	}
	valueData := val.Interface().(string)
	trim := strings.Trim(valueData, " ")
	if len(trim)==0{
		return false,utils.Errorf("匹配值不能为空:%s",valueData)
	}
	return true,nil

}

func (v *validateService)validationCsvType(params map[string]interface{}, val reflect.Value, args ...string)  (bool, error) {
	if val.Kind() != reflect.Int{
		return false,utils.Errorf("csv的文件类型的参数类型错误:%s",val.Kind().String())
	}
	csvType := val.Interface().(int)
	if checkCsvType(csvType)==0{
		return false,utils.Errorf("csv的文件类型不存在:%d",csvType)
	}
	return true,nil
}

// 校验规则id是否存在
func (v *validateService)validationMatchingRuleId(params map[string]interface{}, val reflect.Value, args ...string)  (bool, error)  {
	if val.Kind() != reflect.Int{
		return false,utils.Errorf("规则主键Id的参数类型错误:%s",val.Kind().String())
	}
	id := val.Interface().(int)
	result, _ := v.TypeMatchingRuleDao.FindById(id)
	if checkCsvType(result.CsvType) == 0{
		return false,utils.Errorf("该规则不存在:%d",id)
	}
	return true,nil
}

// 校验匹配模式查询
func (v *validateService)validationFuzzy(params map[string]interface{}, val reflect.Value, args ...string)  (bool, error)  {
	if val.Kind() != reflect.Int{
		return false,utils.Errorf("匹配查询方式的参数类型错误:%s",val.Kind().String())
	}
	id := val.Interface().(int)
	if len(constant.FuzzyNameMap[id]) == 0{
		return false,utils.Errorf("匹配查询方式不存在:%d",id)
	}
	return true,nil
}

func (v *validateService) CheckConsumptionTypeId(consumptionTypeId int) (err error) {
	check, err := v.ConsumptionTypeDao.CheckConsumptionTypeId(consumptionTypeId)
	if err != nil{
		return err
	}
	if !check{
		err = utils.Errorf("钱迹消费类型不存在:%d",consumptionTypeId)
	}
	return
}

func (v *validateService) CheckRulesTypeIdByCsvType(csvType int,rulesTypeId int) (err error) {
	checkBol, err := v.RulesTypeDao.FindRulesCountByCsvType(csvType, rulesTypeId)
	if err != nil{
		return err
	}
	if !checkBol{
		return utils.Errorf("%s的csv表标题不匹配",constant.CsvTypeName[csvType])
	}
	return nil
}


