package routes

import (
	"QianJi_Excel/bootstrap"
	"QianJi_Excel/service"
	"QianJi_Excel/web/controller"
	"QianJi_Excel/web/middleware"
	"fmt"
	"github.com/kataras/iris/mvc"
)
// 路由注册
func RoutesConfigure(b * bootstrap.Bootstrapper)  {
	matchingRuleService := service.NewMatchingRuleService()
	rulesTypeService := service.NewRulesTypeService()
	consumptionTypeService := service.NewConsumptionTypeService()
	validateService := service.NewValidateService()
	// 参数化路径
	sprintf := fmt.Sprintf("/%s", b.APIPrefix)
	application := mvc.New(b.Party(sprintf,middleware.CorsMiddleware()))
	matching_rule := application.Party("/matching_rule")
	// 依赖注入service服务到controller
	matching_rule.Register(matchingRuleService,rulesTypeService,consumptionTypeService,validateService)
	matching_rule.Handle(new(controller.MatchingRuleController))
}
