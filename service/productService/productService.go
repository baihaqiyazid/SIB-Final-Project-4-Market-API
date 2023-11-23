package productservice

import (
	"market-api/models"
	"market-api/web"

	"github.com/gin-gonic/gin"
)

type ProductService interface{
	Create(ctx *gin.Context, request web.CreateProductPayload) (*models.Product, error) 
	GetAll(ctx *gin.Context) (*[]models.Product, error) 
	GetProductById(ctx *gin.Context) (*models.Product, error) 
	Update(ctx *gin.Context, request web.CreateProductPayload) (*models.Product, error) 
	Delete(ctx *gin.Context) (*string, error) 
}