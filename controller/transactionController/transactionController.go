package transactioncontroller

import "github.com/gin-gonic/gin"

type TransactionController interface{
	Create(ctx *gin.Context)
	GetTransactionById(ctx *gin.Context)
	GetAllUserTransaction(ctx *gin.Context)
}