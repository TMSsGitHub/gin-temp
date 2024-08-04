package main

import (
	"fmt"
	"gin-temp/api/middleware"
	"gin-temp/api/route"
	"gin-temp/conf"
	_ "gin-temp/docs"
	"gin-temp/internal/global/datastore"
	"gin-temp/internal/global/logger"
	"github.com/gin-gonic/gin"
)

// main
// @title gin-temp接口管理
// @version 1.0
// @description 简单的接口管理文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apche.org/licenses/LICENSE-2.0.html

// @host localhost:10101
// @BasePath /
// @Schemes http
func main() {
	r := gin.New()
	r.Use(logger.ZapLogger(), gin.Recovery())
	// 注册中间件
	middleware.InitMiddleware(r)
	// 注册路由
	route.InitRouter(r)

	addr := fmt.Sprintf("%s:%d", conf.Cfg.App.Host, conf.Cfg.App.Port)
	err := r.Run(addr)
	if err != nil {
		panic(err)
	}
}

func init() {
	// 读取配置文件
	conf.InitConfig()
	// 初始化日志
	logger.InitLogger()
	// 初始化数据库
	datastore.InitDB()
	// 初始化缓存
	datastore.InitCache()
}
