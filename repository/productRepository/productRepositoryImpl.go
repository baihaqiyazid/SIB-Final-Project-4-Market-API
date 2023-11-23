package productrepository

import (
	"market-api/helper"
	"market-api/models"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

// Create implements ProductRepository.
func (repository *ProductRepositoryImpl) Create(Product models.Product) (*models.Product, error) {
	err := repository.db.Create(&Product).Error
	helper.LogIfError(err)

	return &Product, nil
}

func (repository *ProductRepositoryImpl) GetAll() (*[]models.Product, error) {
	var products []models.Product
	err := repository.db.Find(&products).Error
	helper.LogIfError(err)

	return &products, nil
}

func (repository *ProductRepositoryImpl) Update(id uint, product models.Product) (*models.Product, error) {
    // Menetapkan model yang ingin diupdate dengan menggunakan Model
    err := repository.db.Where("id", id).Updates(&product).Error
	helper.LogIfError(err)

	err = repository.db.First(&product, id).Error
    helper.LogIfError(err)

    return &product, nil
}

func (repository *ProductRepositoryImpl) GetProductById(id uint) (*models.Product, error) {
	var product *models.Product
	err := repository.db.Where("id", id).First(&product).Error
	if err != nil{
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Delete(id uint) (*string, error) {
	var product *models.Product
	err := repository.db.Where("id", id).Delete(&product).Error
	if err != nil{
		return nil, err
	}

	msg := "Product has been successfully deleted"

	return &msg, nil
}