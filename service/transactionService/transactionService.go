package transactionservice

import (
	"market-api/models"
	"market-api/web"

	"github.com/gin-gonic/gin"
)

type TransactionService interface{
	Create(ctx *gin.Context, request web.CreateTransactionPayload) (*models.TransactionHistory, error) 
	GetTransactionById(ctx *gin.Context) (*[]models.TransactionHistory, error) 
	GetAllUserTransaction(ctx *gin.Context) (*[]models.TransactionHistory, error) 
}