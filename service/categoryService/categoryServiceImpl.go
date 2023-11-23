package categoryservice

import (
	"errors"
	"market-api/models"
	repository "market-api/repository/categoryRepository"
	userRepository "market-api/repository/userRepository"
	"market-api/web"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	DB             *gorm.DB
	CategoryRepository repository.CategoryRepository
	UserRepository userRepository.UserRepository

}

func NewCategoryService(db *gorm.DB, categoryRepository repository.CategoryRepository, userRepository userRepository.UserRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		DB:             db,
		CategoryRepository: categoryRepository,
		UserRepository: userRepository,
	}
}

func (service *CategoryServiceImpl) Create(ctx *gin.Context, request web.CreateCategoryPayload) (*models.Category, error) {
	var category models.Category

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
	
	category.Type = request.Type
	data, err := service.CategoryRepository.Create(category);
	if err != nil{
		return nil, errors.New("Email is already registered")
	}

	return data, nil
}

func (service *CategoryServiceImpl) GetAll(ctx *gin.Context) (*[]models.Category, error) {
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
	
	data, err := service.CategoryRepository.GetAll();
	if err != nil{
		return nil, err
	}

	return data, nil
}

func (service *CategoryServiceImpl) GetCategoryById(ctx *gin.Context) (*models.Category, error) {
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

	idParam := ctx.Param("categoryId")

	// Konversi nilai idParam ke tipe data integer
	id, _ := strconv.Atoi(idParam)

	
	data, err := service.CategoryRepository.GetCategoryById(uint(id));
	if err != nil{
		return nil, errors.New("Category not found")
	}

	return data, nil
}

func (service *CategoryServiceImpl) Update(ctx *gin.Context, request web.UpdateCategoryPayload) (*models.Category, error) {
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

	idParam := ctx.Param("categoryId")

	// Konversi nilai idParam ke tipe data integer
	id, _ := strconv.Atoi(idParam)
    data, err := service.CategoryRepository.Update(uint(id), request.Type)
    if err != nil {
        return nil, err
    }

    return data, nil
}

func (service *CategoryServiceImpl) Delete(ctx *gin.Context) (*string, error) {
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

	idParam := ctx.Param("categoryId")

	// Konversi nilai idParam ke tipe data integer
	id, _ := strconv.Atoi(idParam)
    message, err := service.CategoryRepository.Delete(uint(id))
    if err != nil {
        return nil, err
    }

    return message, nil
}