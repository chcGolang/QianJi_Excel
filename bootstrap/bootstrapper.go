package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/view"
	"time"
)

type  Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	// 项目名称
	AppName string
	// 项目负责人
	AppOwner string
	// url前缀
	APIPrefix string
	// 项目启动时间
	AppSpawnDate time.Time
}

// 创建Bootstrapper
func NewBootstrapper(appName,appOwner,apiPrefix string,cfgs ...Configurator)*Bootstrapper  {
	bootstrapper := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    appOwner,
		APIPrefix:   apiPrefix,
		AppSpawnDate:time.Now(),
	}
	for _,cfg:=range cfgs{
		cfg(bootstrapper)
	}
	return bootstrapper
}

const pageSize  = 2
// 初始化html模板
func (b *Bootstrapper)InitView(viewDir string)  {
	var(
		htmlEngine *view.HTMLEngine
	)
	htmlEngine = iris.HTML(viewDir, ".html")
	// 每次重新加载模版（线上关闭它）
	htmlEngine.Reload(true)

	htmlEngine.AddFunc("PreviousPage", func(page int)int {
		if page>0{
			page-=1
		}
		return page

	})

	htmlEngine.AddFunc("NextPage", func(page int)int {
		page+=1
		return page

	})

	htmlEngine.AddFunc("GetPageSize", func()int {

		return pageSize

	})

	b.RegisterView(htmlEngine)
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}


// 加载Bootstrapper的配置
func (b *Bootstrapper)Bootstrap()  {
	// 加载html模板
	b.InitView("./web/views")
	// /js的路径就去查询/web/views/js/下的文件
	b.StaticWeb("/js","./web/views/js/")
	// /css的路径就去查询/web/views/css/下的文件
	b.StaticWeb("/css","./web/views/css/")

	// 设置全局中间件
	b.Use(recover.New())
	b.Use(logger.New())
}

// 错误处理
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("error.html")
	})
}

// 启动服务
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
