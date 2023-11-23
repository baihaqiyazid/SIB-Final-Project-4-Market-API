package userrepository

import "market-api/models"

type UserRepository interface{
	Register(user models.User) (*models.UserResponse, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	TopUp(id uint, balance int) (string, error)
}