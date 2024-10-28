// Package routes 注册路由
package routes

import (
	"GoHub/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，所有 v1 版本的路由都将放在这里
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机号是否已经注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
