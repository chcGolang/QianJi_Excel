package routes

import (
	"QianJi_Excel/bootstrap"
	"QianJi_Excel/service"
	"QianJi_Excel/web/controller"
	"fmt"
	"github.com/kataras/iris/mvc"
)
// 路由注册
func RoutesConfigure(b * bootstrap.Bootstrapper)  {
	matchingRuleService := service.NewMatchingRuleService()
	// 参数化路径
	sprintf := fmt.Sprintf("/%s", b.APIPrefix)
	application := mvc.New(b.Party(sprintf))
	//matching_rule := application.Party("/matching_rule")
	// 依赖注入service服务到controller
	application.Register(matchingRuleService)
	application.Handle(new(controller.MatchingRuleController))
}
