package transactioncontroller

import (
	"log"
	"market-api/helper"
	transactionservice "market-api/service/transactionService"
	"market-api/web"

	"github.com/gin-gonic/gin"
)

type TransactionControllerImpl struct {
	TransactionService transactionservice.TransactionService
}

func NewTransactionController(service transactionservice.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{TransactionService: service}
}

func (controller *TransactionControllerImpl) Create(ctx *gin.Context) {

	var request web.CreateTransactionPayload
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}

	data, err := controller.TransactionService.Create(ctx, request)
	if err != nil{
		if err.Error() == "Product not found" {
			helper.ResponseNotFound(ctx, err.Error())
			return
		}

		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *TransactionControllerImpl) GetTransactionById(ctx *gin.Context) {

	data, err := controller.TransactionService.GetTransactionById(ctx)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}

func (controller *TransactionControllerImpl) GetAllUserTransaction(ctx *gin.Context) {

	data, err := controller.TransactionService.GetAllUserTransaction(ctx)
	if err != nil{
		helper.ResponseBadRequest(ctx, err.Error())
		return
	}
	

	helper.ResponseSuccess(ctx, data)

	return
}