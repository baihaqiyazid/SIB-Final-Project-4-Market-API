package categorycontroller

import "github.com/gin-gonic/gin"

type CategoryController interface{
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}