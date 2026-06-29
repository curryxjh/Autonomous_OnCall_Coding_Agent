package dto

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterRoutes(*gin.Engine)
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
