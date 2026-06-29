package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// 注册自定义密码复杂度验证
	err := validate.RegisterValidation("password_complexity", passwordComplexityValidation)
	if err != nil {
		panic(err)
	}
}

// passwordComplexityValidation 密码复杂度验证
// 要求密码必须同时包含字母和数字
func passwordComplexityValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// 正则表达式：必须同时包含字母和数字
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasLetter && hasDigit
}

// GetValidator 获取验证器实例
func GetValidator() *validator.Validate {
	return validate
}
