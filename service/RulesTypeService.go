package service

import (
	"QianJi_Excel/dao"
	"QianJi_Excel/vo"
)

type IRulesTypeService interface {
	// 根据csv文件类型查询标题
	FindByCsvType(csvType int )(voInfo []vo.RulesTypeInfoVo,err error)
}

func NewRulesTypeService()IRulesTypeService  {
	return &rulesTypeService{
		RulesTypeDao:dao.NewRulesTypeDao(),
	}
}

type rulesTypeService struct {
	RulesTypeDao dao.RulesTypeDao
}

func (r *rulesTypeService) FindByCsvType(csvType int )(voInfo []vo.RulesTypeInfoVo,err error) {
	csvType = checkCsvType(csvType)
	return r.RulesTypeDao.FindByCsvType(csvType)
}

