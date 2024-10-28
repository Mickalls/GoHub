package auth

import (
	v1 "GoHub/app/http/controllers/api/v1"
	"GoHub/app/models/user"
	"GoHub/app/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	request := requests.SignupPhoneExistRequest{}

	// 解析JSON请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败,返回442状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 中断
		return
	}

	// 表单验证
	if errs := requests.ValidateSignupPhoneExist(&request, c); len(errs) > 0 {
		// 验证失败，返回422状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
