package service

import (
	"QianJi_Excel/constant"
	"QianJi_Excel/dao"
	"QianJi_Excel/model"
	"QianJi_Excel/utils"
	"QianJi_Excel/vo"
)

type IRulesTypeService interface {
	// 根据csv文件类型查询标题
	FindByCsvType(csvType int )(voInfo []vo.RulesTypeInfoVo,err error)
	// 分页查询csv标题列表
	FindRulesTypeToPage(titleName string,csv_type int,page,pageSize int)(voInfo []vo.RulesTypeInfoPageVo,count int,err error)
	// 根据Id查询RulesType
	FindRulesTypeInfoById(id int)(voInfo vo.RulesTypeFindVo,err error)
	// 根据Id更新
	UpdateById(id,csv_type int,rules_name string)(err error)
	// 添加csv标题
	SaveRulesType(saveVo vo.RulesTypeSaveVo)(err error)
	// RulesType根据Id删除
	DeleteRulesTypeById(id int)(err error)
}


func NewRulesTypeService()IRulesTypeService  {
	return &rulesTypeService{
		RulesTypeDao:dao.NewRulesTypeDao(),
	}
}

type rulesTypeService struct {
	RulesTypeDao dao.RulesTypeDao
}

func (r *rulesTypeService) DeleteRulesTypeById(id int) (err error) {
	return r.RulesTypeDao.DeleteById(id)
}

func (r *rulesTypeService) SaveRulesType(saveVo vo.RulesTypeSaveVo) (err error) {
	rulesType := model.RulesType{}
	if err = utils.DeepCopy(&rulesType, saveVo);err != nil{
		return
	}
	if err = NewValidateService().CheckRulesName(rulesType.RulesName, rulesType.CsvType);err != nil{
		return
	}
	return r.RulesTypeDao.Save(&rulesType)
}

func (r *rulesTypeService) UpdateById(id,csv_type int, rules_name string)(err error) {
	if err = NewValidateService().CheckRulesName(rules_name, csv_type);err != nil{
		return
	}
	return r.RulesTypeDao.UpdateById(id,rules_name)
}

func (r *rulesTypeService) FindRulesTypeInfoById(id int) (voInfo vo.RulesTypeFindVo, err error) {
	if voInfo, err = r.RulesTypeDao.FindById(id);err != nil{
		return
	}
	voInfo.CsvTypeName = constant.CsvTypeName[voInfo.CsvType]
	return
}

func (r *rulesTypeService) FindRulesTypeToPage(titleName string, csv_type int, page, pageSize int) (voInfo []vo.RulesTypeInfoPageVo,count int, err error) {
	csv_type = checkCsvType(csv_type)
	assemblePage(&page,&pageSize)
	if voInfo, count, err = r.RulesTypeDao.FindToPage(titleName, csv_type, page, pageSize);err != nil{
		return
	}
	for key,value:=range voInfo {
		voInfo[key].CsvTypeName = constant.CsvTypeName[value.CsvType]
	}
	return
}

func (r *rulesTypeService) FindByCsvType(csvType int )(voInfo []vo.RulesTypeInfoVo,err error) {
	csvType = checkCsvType(csvType)
	return r.RulesTypeDao.FindByCsvType(csvType)
}

