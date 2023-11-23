package categoryrepository

import "market-api/models"

type CategoryRepository interface{
	Create(category models.Category) (*models.Category, error)
	GetAll() (*[]models.Category, error)
	GetCategoryById(id uint) (*models.Category, error)
	Delete(id uint) (*string, error)
	Update(id uint, name string) (*models.Category, error)
}