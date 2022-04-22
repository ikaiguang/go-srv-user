package main

import (
	stdlog "log"
	"os"

	"github.com/go-kratos/kratos/v2"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	routes "github.com/ikaiguang/go-srv-user/internal/route"
	servers "github.com/ikaiguang/go-srv-user/internal/server"
	"github.com/ikaiguang/go-srv-user/internal/setup"
)

// newApp .
func newApp(modulesHandler setup.Modules) (app *kratos.App, err error) {
	// 主机
	hostname, _ := os.Hostname()

	// 日志
	logger, _, err := modulesHandler.Logger()
	if err != nil {
		return app, err
	}

	// 服务
	hs, err := servers.NewHTTPServer(modulesHandler)
	if err != nil {
		return app, err
	}
	gs, err := servers.NewGRPCServer(modulesHandler)
	if err != nil {
		return app, err
	}

	// 路由
	err = routes.RegisterRoutes(modulesHandler, hs, gs)
	if err != nil {
		return app, err
	}

	// app
	appConfig := modulesHandler.AppConfig()
	app = kratos.New(
		kratos.ID(hostname),
		kratos.Name(appConfig.Name),
		kratos.Version(appConfig.Version),
		kratos.Metadata(appConfig.Metadata),
		kratos.Logger(logger),
		kratos.Server(hs, gs),
	)
	return app, err
}

// go run ./example/cmd/main/... -conf=./example/configs
func main() {
	var err error

	// 初始化
	err = setup.Init()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	//defer func() { _ = setup.Close() }()

	// 模块
	modules, err := setup.GetModules()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	defer func() { _ = setup.Close() }()

	// 启动程序
	stdlog.Println()
	stdlog.Println("|==================== 启动程序 开始 ====================|")

	// 启动
	app, err := newApp(modules)
	stdlog.Println("|==================== 启动程序 结束 ====================|")
	if err := app.Run(); err != nil {
		debugutil.Fatalf("%+v\n", err)
		return
	}
}
