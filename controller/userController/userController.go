package usercontroller

import "github.com/gin-gonic/gin"

type UserController interface{
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	TopUp(ctx *gin.Context)
}