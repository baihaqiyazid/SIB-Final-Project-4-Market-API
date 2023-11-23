package transactionservice

import (
	"errors"
	"fmt"
	"market-api/models"
	productrepository "market-api/repository/productRepository"
	transactionrepository "market-api/repository/transactionRepository"
	userrepository "market-api/repository/userRepository"
	"market-api/web"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	DB             *gorm.DB
	TransactionRepository transactionrepository.TransactionRepository
	UserRepository userrepository.UserRepository
	ProductRepository productrepository.ProductRepository
}

func NewTransactionService(db *gorm.DB, 
	transactionRepository transactionrepository.TransactionRepository, 
	userRepository userrepository.UserRepository,
	productRepository productrepository.ProductRepository,
	) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		DB:             db,
		TransactionRepository: transactionRepository,
		UserRepository: userRepository,
		ProductRepository: productRepository,
	}
}

func (service *TransactionServiceImpl) Create(ctx *gin.Context, request web.CreateTransactionPayload) (*models.TransactionHistory, error) {
	var transaction models.TransactionHistory

	userIDFloat, exists := ctx.Get("userID")
    if !exists {
        // Jika tidak ada, berikan respons error atau lakukan sesuatu sesuai kebutuhan
        return nil, errors.New("Internal Server Error")
    }

	// Konversi nilai userID ke tipe data int
    userID := uint(userIDFloat.(float64))
	user, _ := service.UserRepository.GetUserById(userID)

	product, err := service.ProductRepository.GetProductById(request.ProductID)
	fmt.Println("user.Balance: ", user.Balance)
	fmt.Println("product.Price: ", product.Price)
	fmt.Println("product.Price: ", product.Price)
	fmt.Println("request.Quantity: ", request.Quantity)
	fmt.Println("product.Price * request.Quantity: ", product.Price * request.Quantity)
	fmt.Println("user.Balance - (product.Price * request.Quantity): ", user.Balance - (product.Price * request.Quantity))
	fmt.Println("user.Balance < (user.Balance - (product.Price * request.Quantity)): ", user.Balance < (user.Balance - (product.Price * request.Quantity)))

	if err != nil{
		return nil, errors.New("Product not found")
	}

	if(request.Quantity > product.Stock){
		return nil, errors.New("Quantity must be less")
	}

	if((user.Balance - (product.Price * request.Quantity) < 0)){
		return nil, errors.New("Your balance is not enough")
	}

	newBalance := user.Balance - (product.Price * request.Quantity)
	newStock := product.Stock - request.Quantity

	err = service.TransactionRepository.UpdateStockProduct(product.ID, newStock)
	if err != nil{
		return nil, err
	}

	err = service.TransactionRepository.UpdateBalanceUser(userID, newBalance)
	if err != nil{
		return nil, err
	}

	err = service.TransactionRepository.UpdateSoldAmountCategory(product.CategoryID, request.Quantity)
	if err != nil{
		return nil, err
	}

	transaction.ProductID = request.ProductID
	transaction.Quantity = uint(request.Quantity)
	transaction.UserID = userID
	transaction.TotalPrice = product.Price * request.Quantity
	transaction.UpdatedAt = time.Now()

	data, err := service.TransactionRepository.Create(transaction);
	if err != nil{
		return nil, err
	}

	return data, nil
}

func (service *TransactionServiceImpl) GetTransactionById(ctx *gin.Context) (*[]models.TransactionHistory, error) {
	userIDFloat, exists := ctx.Get("userID")
    if !exists {
        // Jika tidak ada, berikan respons error atau lakukan sesuatu sesuai kebutuhan
        return nil, errors.New("Internal Server Error")
    }

    // Konversi nilai userID ke tipe data int
    userID := uint(userIDFloat.(float64))
	
	data, err := service.TransactionRepository.GetTransactionById(userID);
	if err != nil{
		return nil, errors.New("Transaction not found")
	}

	return data, nil
}

func (service *TransactionServiceImpl) GetAllUserTransaction(ctx *gin.Context) (*[]models.TransactionHistory, error) {
	userIDFloat, exists := ctx.Get("userID")
    if !exists {
        // Jika tidak ada, berikan respons error atau lakukan sesuatu sesuai kebutuhan
        return nil, errors.New("Internal Server Error")
    }

    // Konversi nilai userID ke tipe data int
    userID := uint(userIDFloat.(float64))

	user, _ := service.UserRepository.GetUserById(userID)
	if user.Role != "admin" {
		return nil, errors.New("Your account cannot access this feature!")
	}
	
	data, err := service.TransactionRepository.GetAllUserTransaction();
	if err != nil{
		return nil, errors.New("Transaction not found")
	}

	return data, nil
}