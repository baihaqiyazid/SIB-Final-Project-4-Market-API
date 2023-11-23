package userrepository

import (
	"market-api/helper"
	"market-api/models"
	"strconv"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repository *UserRepositoryImpl) Register(user models.User) (*models.UserResponse, error) {
	err := repository.db.Create(&user).Error
	helper.LogIfError(err)

	return &models.UserResponse{
		ID: user.ID,
		Email: user.Email,
		Fullname: user.Fullname,
		Balance: user.Balance,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (repository *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := repository.db.Where("email", email).First(&user).Error
	helper.LogIfError(err)

	return user, nil
}

func (repository *UserRepositoryImpl) GetUserById(id uint) (*models.User, error) {
	var user *models.User
	err := repository.db.Where("id", id).First(&user).Error
	helper.LogIfError(err)

	return user, nil
}

func (repository *UserRepositoryImpl) TopUp(id uint, balance int) (string, error) {
    // Membuat instansiasi model User dengan ID yang sesuai
    user := models.User{ID: id}

    // Menetapkan model yang ingin diupdate dengan menggunakan Model
    err := repository.db.Model(&user).UpdateColumn("balance", balance).Error
    helper.LogIfError(err)

    msg := "Your balance successfully updated to Rp " + strconv.Itoa(balance)

    return msg, nil
}

