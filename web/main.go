package main

import (
	"QianJi_Excel/bootstrap"
	"QianJi_Excel/web/routes"
	"fmt"
)

func main()  {
	bootstrapper := bootstrap.NewBootstrapper("钱迹模板转换", "chc", "qianji")
	bootstrapper.Bootstrap()
	bootstrapper.Configure(routes.RoutesConfigure)
	bootstrapper.Listen(fmt.Sprintf(":%d",8080))
}
