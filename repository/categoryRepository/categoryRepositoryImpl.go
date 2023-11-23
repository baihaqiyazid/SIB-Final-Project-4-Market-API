package categoryrepository

import (
	"market-api/helper"
	"market-api/models"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

// Create implements CategoryRepository.
func (repository *CategoryRepositoryImpl) Create(category models.Category) (*models.Category, error) {
	err := repository.db.Create(&category).Error
	helper.LogIfError(err)

	return &category, nil
}

func (repository *CategoryRepositoryImpl) GetAll() (*[]models.Category, error) {
	var categories []models.Category
	err := repository.db.Preload("Products").Find(&categories).Error
	helper.LogIfError(err)

	return &categories, nil
}

func (repository *CategoryRepositoryImpl) Update(id uint, name string) (*models.Category, error) {
    // Membuat instansiasi model User dengan ID yang sesuai
    category := models.Category{ID: id}

    // Menetapkan model yang ingin diupdate dengan menggunakan Model
    err := repository.db.Model(&category).UpdateColumn("type", name).Error
	helper.LogIfError(err)

	err = repository.db.Preload("Products").First(&category, id).Error
    helper.LogIfError(err)

    return &category, nil
}

func (repository *CategoryRepositoryImpl) GetCategoryById(id uint) (*models.Category, error) {
	var category *models.Category
	err := repository.db.Where("id", id).First(&category).Error
	if err != nil{
		return nil, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Delete(id uint) (*string, error) {
	var category *models.Category
	err := repository.db.Where("id", id).Delete(&category).Error
	if err != nil{
		return nil, err
	}

	msg := "Category has been successfully deleted"

	return &msg, nil
}