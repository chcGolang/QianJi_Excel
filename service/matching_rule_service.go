package service

import (
	"QianJi_Excel/dao"
	"QianJi_Excel/vo"
)

// 规则操作服务
type IMatchingRuleService interface {
	// 查询匹配规则列表
	FindTypeMatchingRuleBypage(csvType int, page,pageSize int)(resultList []vo.TypeMatchingRuleVo,count int,err error)
}
func NewMatchingRuleService() IMatchingRuleService {
	return &matchingRuleService{
		typeMatchingRuleDao:dao.NewTypeMatchingRuleDao(),
	}
}
type matchingRuleService struct {
	typeMatchingRuleDao dao.TypeMatchingRuleDao
}

func (v *matchingRuleService) FindTypeMatchingRuleBypage(csvType int, page,pageSize int)(resultList []vo.TypeMatchingRuleVo,count int,err error) {
	assemblePage(&page,&pageSize)
	return v.typeMatchingRuleDao.FindTypeMatchingRule(csvType,page,pageSize)
}

