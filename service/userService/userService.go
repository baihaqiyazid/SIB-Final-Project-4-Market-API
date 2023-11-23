package userservice

import (
	"market-api/models"
	"market-api/web"

	"github.com/gin-gonic/gin"
)

type UserService interface{
	Register(ctx *gin.Context, request web.RegisterPayload) (*models.UserResponse, error) 
	Login(ctx *gin.Context, request web.LoginPayload) (*string, error) 
	TopUp(ctx *gin.Context, request web.TopupPayload) (*string, error) 
}