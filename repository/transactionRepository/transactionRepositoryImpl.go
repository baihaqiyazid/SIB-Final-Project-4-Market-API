package transactionrepository

import (
	"market-api/models"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

// Create implements TransactionRepository.
func (repository *TransactionRepositoryImpl) Create(transaction models.TransactionHistory) (*models.TransactionHistory, error) {
	// product := models.Product{ID: transaction.ProductID}
	
	err := repository.db.Create(&transaction).Error
	if err != nil{
		return nil, err
	 }

	err = repository.db.Preload("Product").First(&transaction).Error
	if err != nil{
		return nil, err
	 }
	return &transaction, nil
}

func (repository *TransactionRepositoryImpl) UpdateBalanceUser(id uint, balance int) error {
	 // Membuat instansiasi model User dengan ID yang sesuai
	 user := models.User{ID: id}

	 // Menetapkan model yang ingin diupdate dengan menggunakan Model
	 err := repository.db.Model(&user).UpdateColumn("balance", balance).Error
	 if err != nil{
		return err
	 }
 
	 return nil
}

func (repository *TransactionRepositoryImpl) UpdateStockProduct(id uint, stock int) error {
	// Membuat instansiasi model User dengan ID yang sesuai
	product := models.Product{ID: id}

	// Menetapkan model yang ingin diupdate dengan menggunakan Model
	err := repository.db.Model(&product).UpdateColumn("stock", stock).Error
	if err != nil{
	   return err
	}

	return nil
}

func (repository *TransactionRepositoryImpl) UpdateSoldAmountCategory(id uint, amount int) error {
	category := models.Category{ID: id}

	// Menetapkan model yang ingin diupdate dengan menggunakan Model
	err := repository.db.Model(&category).UpdateColumn("sold_product_amount", amount).Error
	if err != nil{
	   return err
	}

	return nil
}

func (repository *TransactionRepositoryImpl) GetTransactionById(id uint) (*[]models.TransactionHistory, error) {
	var transaction *[]models.TransactionHistory
	err := repository.db.Where("user_id", id).Preload("Product").Find(&transaction).Error
	if err != nil{
		return nil, err
	}

	return transaction, nil
}

func (repository *TransactionRepositoryImpl) GetAllUserTransaction() (*[]models.TransactionHistory, error) {
	var transaction *[]models.TransactionHistory
	err := repository.db.Preload("Product").Preload("User").Find(&transaction).Error
	if err != nil{
		return nil, err
	}

	return transaction, nil
}