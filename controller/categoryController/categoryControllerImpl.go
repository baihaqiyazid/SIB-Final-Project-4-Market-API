package categorycontroller

import (
	"log"
	"market-api/helper"
	categoryservice "market-api/service/categoryService"
	"market-api/web"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService categoryservice.CategoryService
}

func NewCategoryController(service categoryservice.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{CategoryService: service}
}

func (controller *CategoryControllerImpl) Create(ctx *gin.Context) {

	var request web.CreateCategoryPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.CategoryService.Create(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *CategoryControllerImpl) GetAll(ctx *gin.Context) {

	data, err := controller.CategoryService.GetAll(ctx)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *CategoryControllerImpl) Update(ctx *gin.Context) {

	var request web.UpdateCategoryPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	_, err := controller.CategoryService.GetCategoryById(ctx)
	if err != nil{
		helper.ResponseNotFound(ctx, err.Error())
		return
	}

	data, err := controller.CategoryService.Update(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *CategoryControllerImpl) Delete(ctx *gin.Context) {

	_, err := controller.CategoryService.GetCategoryById(ctx)
	if err != nil{
		helper.ResponseNotFound(ctx, err.Error())
		return
	}

	data, err := controller.CategoryService.Delete(ctx)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseCategoryMessage(ctx, data)

	return
}

