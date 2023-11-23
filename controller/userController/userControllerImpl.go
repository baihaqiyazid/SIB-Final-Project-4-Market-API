package usercontroller

import (
	"market-api/helper"
	service "market-api/service/userService"
	"market-api/web"
	"log"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: service}
}

func (controller *UserControllerImpl) Register(ctx *gin.Context) {

	var request web.RegisterPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.UserService.Register(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *UserControllerImpl) Login(ctx *gin.Context) {

	var request web.LoginPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.UserService.Login(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseSuccessLogin(ctx, data)

	return
}

func (controller *UserControllerImpl) TopUp(ctx *gin.Context) {

	var request web.TopupPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.UserService.TopUp(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseSuccessTopup(ctx, data)

	return
}