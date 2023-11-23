package productcontroller

import "github.com/gin-gonic/gin"

type ProductController interface{
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}