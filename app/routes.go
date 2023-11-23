package app

import (
	"fmt"
	categorycontroller "market-api/controller/categoryController"
	productcontroller "market-api/controller/productController"
	transactioncontroller "market-api/controller/transactionController"
	userController "market-api/controller/userController"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func Route(
	userController userController.UserController, 
	categoryContoller categorycontroller.CategoryController, 
	productController productcontroller.ProductController,
	transactionController transactioncontroller.TransactionController,
	){
	router := gin.Default()
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    router.Use(cors.New(config))

    // Grup rute yang memerlukan otentikasi dengan JWT
    authGroup := router.Group("/users")
	categoryGroup := router.Group("/categories")
	productGroup := router.Group("/products")
	transactionGroup := router.Group("/transactions")


	authGroup.POST("/register", userController.Register)
	authGroup.POST("/login", userController.Login)

    authGroup.Use(JWTMiddleware())
    {
        authGroup.PATCH("/topup", userController.TopUp)
    }

	categoryGroup.Use(JWTMiddleware())
	{
		categoryGroup.POST("/", categoryContoller.Create)
		categoryGroup.GET("/", categoryContoller.GetAll)
		categoryGroup.PATCH("/:categoryId", categoryContoller.Update)
		categoryGroup.DELETE("/:categoryId", categoryContoller.Delete)
	}

	productGroup.Use(JWTMiddleware())
	{
		productGroup.POST("/", productController.Create)
		productGroup.GET("/", productController.GetAll)
		productGroup.PUT("/:productId", productController.Update)
		productGroup.DELETE("/:productId", productController.Delete)
	}

	transactionGroup.Use(JWTMiddleware())
	{
		transactionGroup.POST("/", transactionController.Create)
		transactionGroup.GET("/my-transactions", transactionController.GetTransactionById)
		transactionGroup.GET("/user-transactions", transactionController.GetAllUserTransaction)
	}

	router.Run(":4001")
}

// JWTMiddleware adalah middleware untuk mengekstrak dan memverifikasi token JWT dari header Authorization
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Dapatkan token dari header Authorization
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secretKey := []byte(os.Getenv("SECRET_KEY"))
			return secretKey, nil // Ganti dengan secret key yang digunakan untuk menghasilkan token
		})
		
		if err != nil {
			fmt.Println("Error parsing token:", err) // Tambahkan log pesan kesalahan
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Cek apakah token valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Token valid, ambil informasi pengguna dari token
			userID := claims["sub"].(float64) // Ubah sesuai dengan field yang sesuai dalam token
			// Lakukan sesuatu dengan userID, misalnya ambil data pengguna dari database

			// Simpan informasi pengguna di context
			ctx.Set("userID", userID)

			// Lanjutkan ke handler utama
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
	}
}