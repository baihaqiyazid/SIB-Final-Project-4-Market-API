package productservice

import (
	"errors"
	"market-api/models"
	productrepository "market-api/repository/productRepository"
	userrepository "market-api/repository/userRepository"
	"market-api/web"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB             *gorm.DB
	ProductRepository productrepository.ProductRepository
	UserRepository userrepository.UserRepository
}

func NewProductService(db *gorm.DB, productRepository productrepository.ProductRepository, userRepository userrepository.UserRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		DB:             db,
		ProductRepository: productRepository,
		UserRepository: userRepository,
	}
}

func (service *ProductServiceImpl) Create(ctx *gin.Context, request web.CreateProductPayload) (*models.Product, error) {
	var Product models.Product

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
	
	Product.Title = request.Title
	Product.Price = request.Price
	Product.Stock = request.Stock
	Product.CategoryID = request.CategoryId

	data, err := service.ProductRepository.Create(Product);
	if err != nil{
		return nil, err
	}

	return data, nil
}

func (service *ProductServiceImpl) GetAll(ctx *gin.Context) (*[]models.Product, error) {
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
	
	data, err := service.ProductRepository.GetAll();
	if err != nil{
		return nil, err
	}

	return data, nil
}

func (service *ProductServiceImpl) GetProductById(ctx *gin.Context) (*models.Product, error) {
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

	idParam := ctx.Param("productId")

	// Konversi nilai idParam ke tipe data integer
	id, _ := strconv.Atoi(idParam)

	
	data, err := service.ProductRepository.GetProductById(uint(id));
	if err != nil{
		return nil, errors.New("Product not found")
	}

	return data, nil
}

func (service *ProductServiceImpl) Update(ctx *gin.Context, request web.CreateProductPayload) (*models.Product, error) {
    // Mengambil nilai userID dari context
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

	idParam := ctx.Param("productId")

	// Konversi nilai idParam ke tipe data integer
	var product models.Product
	product.Price = request.Price
	product.Title = request.Title
	product.Stock = request.Stock
	product.CategoryID = request.CategoryId
	product.UpdatedAt = time.Now()

	id, _ := strconv.Atoi(idParam)
    data, err := service.ProductRepository.Update(uint(id), product)
    if err != nil {
        return nil, err
    }

    return data, nil
}

func (service *ProductServiceImpl) Delete(ctx *gin.Context) (*string, error) {
    // Mengambil nilai userID dari context
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

	idParam := ctx.Param("productId")

	// Konversi nilai idParam ke tipe data integer
	id, _ := strconv.Atoi(idParam)
    message, err := service.ProductRepository.Delete(uint(id))
    if err != nil {
        return nil, err
    }

    return message, nil
}