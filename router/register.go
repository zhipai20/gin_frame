package router

import (
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"kang/pkg/response"

	"kang/global"
	"kang/middleware"

	"net/http"
)

func Register() *gin.Engine {
	gin.SetMode(global.G_Conf.Server.Mode)
	router := gin.New()

	router.StaticFS(global.G_Conf.Local.Path, http.Dir(global.G_Conf.Local.Path)) // 为用户头像和文件提供静态地址
	//router.Use(middleware.LoadTls())  // 打开就能玩https了

	// 跨域，如需跨域可以打开下面的注释
	router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	// header add X-Request-Id
	router.Use(requestid.New())
	router.Use(middleware.RequestIdAuth())

	//页面找不到
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		response.NotFoundException(c, fmt.Sprintf("%s %s not found", method, path))
	})


	return router
}
