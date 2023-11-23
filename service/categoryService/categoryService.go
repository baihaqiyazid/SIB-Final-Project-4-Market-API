package categoryservice

import (
	"market-api/models"
	"market-api/web"

	"github.com/gin-gonic/gin"
)

type CategoryService interface{
	Create(ctx *gin.Context, request web.CreateCategoryPayload) (*models.Category, error) 
	GetAll(ctx *gin.Context) (*[]models.Category, error) 
	GetCategoryById(ctx *gin.Context) (*models.Category, error) 
	Update(ctx *gin.Context, request web.UpdateCategoryPayload) (*models.Category, error) 
	Delete(ctx *gin.Context) (*string, error) 
}