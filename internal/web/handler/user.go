package handler

import (
	"Autonomous_OnCall_Coding_Agent/internal/web/dto"

	"github.com/gin-gonic/gin"
)

var _ dto.Handler = (*UserHandler)(nil)

type UserHandler struct {
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/user")

	ug.POST("/login", u.Login)
	ug.POST("/signup", u.Signup)
	ug.POST("/edit", u.Edit)
	ug.POST("/logout", u.Logout)
	ug.POST("/refresh_token", u.RefreshToken)

}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Signup(c *gin.Context) {

}

func (u *UserHandler) Login(c *gin.Context) {}

func (u *UserHandler) Edit(c *gin.Context) {}

func (u *UserHandler) Logout(c *gin.Context) {}

func (u *UserHandler) RefreshToken(c *gin.Context) {}
