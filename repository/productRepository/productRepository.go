package productrepository

import "market-api/models"

type ProductRepository interface{
	Create(product models.Product) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	GetProductById(id uint) (*models.Product, error)
	Delete(id uint) (*string, error)
	Update(id uint, product models.Product) (*models.Product, error)
}