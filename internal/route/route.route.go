package routes

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingroute "github.com/ikaiguang/go-srv-user/internal/route/ping"
	testdataroute "github.com/ikaiguang/go-srv-user/internal/route/testdata"
	"github.com/ikaiguang/go-srv-user/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(modulesHandler setup.Modules, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// 日志
	logger, _, err := modulesHandler.Logger()
	if err != nil {
		return err
	}

	stdlog.Println("|*** 注册路由：ping")
	pingroute.RegisterRoutes(hs, gs, logger)

	stdlog.Println("|*** 注册路由：testdata")
	testdataroute.RegisterRoutes(hs, gs, logger)

	return err
}