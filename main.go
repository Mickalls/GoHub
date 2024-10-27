package main

import (
	"GoHub/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 gin engine 实例
	router := gin.New()
	// 初始化绑定路由
	bootstrap.SetUpRoute(router)
	// 运行服务
	if err := router.Run(":3000"); err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
