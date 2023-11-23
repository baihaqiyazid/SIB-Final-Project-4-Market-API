package transactionrepository

import "market-api/models"

type TransactionRepository interface {
	Create(transaction models.TransactionHistory) (*models.TransactionHistory, error)
	UpdateBalanceUser(userId uint, newBalance int) error
	UpdateStockProduct(productId uint, newStock int) error
	UpdateSoldAmountCategory(categoryId uint, newAmount int) error
	GetTransactionById(id uint) (*[]models.TransactionHistory, error)
	GetAllUserTransaction() (*[]models.TransactionHistory, error)
}