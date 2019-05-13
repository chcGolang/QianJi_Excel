package routes

import (
	"QianJi_Excel/bootstrap"
	"QianJi_Excel/service"
	"QianJi_Excel/web/controller"
	"QianJi_Excel/web/middleware"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)
// 路由注册
func RoutesConfigure(b * bootstrap.Bootstrapper)  {
	matchingRuleService := service.NewMatchingRuleService()
	rulesTypeService := service.NewRulesTypeService()
	consumptionTypeService := service.NewConsumptionTypeService()
	validateService := service.NewValidateService()
	csvFileInfoService := service.NewCsvFileInfoService()
	// 参数化路径
	sprintf := fmt.Sprintf("/%s", b.APIPrefix)
	application := mvc.New(b.Party(sprintf,middleware.CorsMiddleware()).AllowMethods(iris.MethodOptions))
	// 依赖注入service服务到controller
	application.Register(matchingRuleService,rulesTypeService,consumptionTypeService,
		validateService,csvFileInfoService)

	// 规则管理
	matching_rule := application.Party("/matching_rule")
	matching_rule.Handle(new(controller.MatchingRuleController))

	// csv标题管理
	csv_title := application.Party("/csv_title")
	csv_title.Handle(new(controller.CsvTitleController))

	// 钱迹消费类型管理
	consumption_type := application.Party("/consumption_type")
	consumption_type.Handle(new(controller.ConsumptionTypeController))

	// csv文件管理
	csv_file := application.Party("/csv_file")
	csv_file.Handle(new(controller.CsvFileController))

}
