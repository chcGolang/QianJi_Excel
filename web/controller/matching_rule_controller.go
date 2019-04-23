package controller

import (
	"QianJi_Excel/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type MatchingRuleController struct {
	Ctx iris.Context
	MatchingRuleService service.IMatchingRuleService
}

func (m *MatchingRuleController)GetList_html()mvc.Result  {
	return mvc.View{
		Name:"index.html",
	}
}

func (m *MatchingRuleController)GetListdata() mvc.Result {
	csv_type, err := m.Ctx.URLParamInt("csv_type")
	page,_:=m.Ctx.URLParamInt("page")
	pageSize,_:=m.Ctx.URLParamInt("pageSize")
	if err != nil{
		return nil
	}
	ckeckPage(&page,&pageSize)
	resultList,count, err := m.MatchingRuleService.FindTypeMatchingRuleBypage(csv_type,page,pageSize)
	if err != nil{
		return nil
	}
	return mvc.View{
		Name:"index.html",
		Data:iris.Map{
			IRIS_DATA:resultList,
			IRIS_TOTAL:count,
			IRIS_PAGE:page,
		},
	}

}
