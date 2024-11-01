// Package bootstrap 处理程序初始化逻辑
package bootstrap

/*
bootstrap包负责程序初始化
暂时只有route.go，后续还会加上MySQL, redis, config....
*/

import (
	"GoHub/app/http/middlewares"
	"GoHub/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetUpRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)
	// 注册 API 路由
	routes.RegisterAPIRoutes(router)
	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回: 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认url和请求方法是否正确",
			})
		}

	})
}
