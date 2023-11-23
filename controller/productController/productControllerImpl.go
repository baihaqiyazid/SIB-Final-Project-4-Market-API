package productcontroller

import (
	"log"
	"market-api/helper"
	productservice "market-api/service/productService"
	"market-api/web"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService productservice.ProductService
}

func NewProductController(service productservice.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{ProductService: service}
}

func (controller *ProductControllerImpl) Create(ctx *gin.Context) {

	var request web.CreateProductPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.ProductService.Create(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err)
		return
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *ProductControllerImpl) GetAll(ctx *gin.Context) {

	data, err := controller.ProductService.GetAll(ctx)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *ProductControllerImpl) Update(ctx *gin.Context) {

	var request web.CreateProductPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	_, err := controller.ProductService.GetProductById(ctx)
	if err != nil{
		helper.ResponseNotFound(ctx, err.Error())
		return
	}

	data, err := controller.ProductService.Update(ctx, request)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *ProductControllerImpl) Delete(ctx *gin.Context) {

	_, err := controller.ProductService.GetProductById(ctx)
	if err != nil{
		helper.ResponseNotFound(ctx, err.Error())
		return
	}

	data, err := controller.ProductService.Delete(ctx)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	
	helper.ResponseCategoryMessage(ctx, data)

	return
}

