package userservice

import (
	"errors"
	"market-api/helper"
	"market-api/models"
	repository "market-api/repository/userRepository"
	"market-api/web"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
}

func NewUserService(db *gorm.DB, userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		DB:             db,
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Register(ctx *gin.Context, request web.RegisterPayload) (*models.UserResponse, error) {
	var user models.User

	_, err := service.UserRepository.GetUserByEmail(request.Email);
	if err != nil{
		return nil, errors.New("Email is already registered")
	}

	user.Fullname = request.Fullname
	user.Email = request.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.Balance = 0
	user.Role = "customer"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	data, err := service.UserRepository.Register(user)
	helper.LogIfError(err)

	return data, nil
}

// Login implements the login functionality with JWT authentication.
func (service *UserServiceImpl) Login(ctx *gin.Context, request web.LoginPayload) (*string, error) {
	// Get user by email
	user, err := service.UserRepository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New("Incorrect email or password")
	}

	// Check password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("Incorrect email or password")
	}

	// Generate JWT token
	token, err := generateJWTToken(*user)
	if err != nil {
		return nil, errors.New("Failed to generate JWT token")
	}

	// Set token in header
	ctx.Header("Authorization", "Bearer "+token)

	return &token, nil
}

// generateJWTToken generates a JWT token for the user.
func generateJWTToken(user models.User) (string, error) {
	// Define the token claims
	claims := jwt.MapClaims{
		"sub": user.ID,
		"iss": "your-issuer", // Set your own issuer
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY"))) // Set your own secret key
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *UserServiceImpl) TopUp(ctx *gin.Context, request web.TopupPayload) (*string, error) {
    // Mengambil nilai userID dari context
    userIDFloat, exists := ctx.Get("userID")
    if !exists {
        // Jika tidak ada, berikan respons error atau lakukan sesuatu sesuai kebutuhan
        return nil, errors.New("Internal Server Error")
    }

    // Konversi nilai userID ke tipe data int
    userID := uint(userIDFloat.(float64))
	balance := request.Balance

	if balance <= 0 {
		return nil, errors.New("Balance must be rather than Rp 0")
	}else if balance >= 100000000{
		return nil, errors.New("Balance must be less than Rp 100.000.000")
	}
	
	// Update Balance
	user, err := service.UserRepository.GetUserById(userID)
	balance = balance + user.Balance

    msg, err := service.UserRepository.TopUp(userID, balance)
    if err != nil {
        return nil, err
    }

    return &msg, nil
}

