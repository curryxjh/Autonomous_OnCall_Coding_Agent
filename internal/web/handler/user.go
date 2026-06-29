package handler

import (
	"Autonomous_OnCall_Coding_Agent/internal/domain/user"
	"Autonomous_OnCall_Coding_Agent/internal/service"
	"Autonomous_OnCall_Coding_Agent/internal/web/dto"
	"Autonomous_OnCall_Coding_Agent/internal/web/middleware"
	ijwt "Autonomous_OnCall_Coding_Agent/internal/web/jwt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

var _ dto.Handler = (*UserHandler)(nil)

type UserHandler struct {
	svc service.UserService
	cmd redis.Cmdable
	ijwt.Handler
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/user")

	ug.POST("/login", middleware.ValidationMiddleware[dto.LoginReq](), u.Login)
	ug.POST("/signup", middleware.ValidationMiddleware[dto.SignupReq](), u.Signup)
	ug.POST("/edit", u.Edit)
	ug.POST("/logout", u.Logout)
	ug.POST("/refresh_token", u.RefreshToken)
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Signup(ctx *gin.Context) {
	// 从中间件获取已验证的数据
	validatedData, exists := ctx.Get("validated_data")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证数据不存在"})
		return
	}
	
	req := validatedData.(dto.SignupReq)

	err := u.svc.SignUp(ctx.Request.Context(), user.User{Email: req.Email, Username: req.Username, Password: req.Password})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func (u *UserHandler) Login(ctx *gin.Context) {
	// 从中间件获取已验证的数据
	validatedData, exists := ctx.Get("validated_data")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证数据不存在"})
		return
	}
	
	req := validatedData.(dto.LoginReq)
	
	// TODO: 实现登录逻辑
	ctx.JSON(http.StatusOK, gin.H{"message": "登录成功", "email": req.Email})
}

func (u *UserHandler) Edit(c *gin.Context) {}

func (u *UserHandler) Logout(c *gin.Context) {}

func (u *UserHandler) RefreshToken(c *gin.Context) {}