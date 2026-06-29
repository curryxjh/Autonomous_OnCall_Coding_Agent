package middleware

import (
	custom_validator "Autonomous_OnCall_Coding_Agent/internal/web/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// ValidationMiddleware 通用的验证中间件
func ValidationMiddleware[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		
		// 绑定JSON数据
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
			c.Abort()
			return
		}
		
		// 使用自定义验证器进行验证
		validate := custom_validator.GetValidator()
		if err := validate.Struct(req); err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				var errorMessages []string
				for _, fieldError := range validationErrors {
					errorMessages = append(errorMessages, getErrorMessage(fieldError))
				}
				c.JSON(http.StatusBadRequest, gin.H{"error": strings.Join(errorMessages, ", ")})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}
		
		// 将验证通过的数据存入上下文
		c.Set("validated_data", req)
		c.Next()
	}
}

// getErrorMessage 根据验证错误生成错误消息
func getErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Field() {
	case "Username":
		switch fieldError.Tag() {
		case "required":
			return "用户名不能为空"
		case "min", "max":
			return "用户名长度必须在1-64个字符之间"
		}
	case "Email":
		switch fieldError.Tag() {
		case "required":
			return "邮箱不能为空"
		case "email":
			return "邮箱格式不正确"
		case "max":
			return "邮箱长度不能超过128个字符"
		}
	case "Password":
		switch fieldError.Tag() {
		case "required":
			return "密码不能为空"
		case "min":
			return "密码长度不能少于6位"
		case "max":
			return "密码长度不能超过255个字符"
		case "password_complexity":
			return "密码必须同时包含字母和数字"
		}
	case "ConfirmPassword":
		switch fieldError.Tag() {
		case "required":
			return "确认密码不能为空"
		case "eqfield":
			return "两次输入的密码不一致"
		}
	}
	
	return fieldError.Error()
}