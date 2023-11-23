package main

import (
	"log"
	"market-api/app"
	categorycontroller "market-api/controller/categoryController"
	productcontroller "market-api/controller/productController"
	transactioncontroller "market-api/controller/transactionController"
	userController "market-api/controller/userController"
	transactionrepository "market-api/repository/transactionRepository"
	categoryrepository "market-api/repository/categoryRepository"
	productrepository "market-api/repository/productRepository"
	userRepository "market-api/repository/userRepository"
	transactionservice "market-api/service/transactionService"
	categoryservice "market-api/service/categoryService"
	productservice "market-api/service/productService"
	userService "market-api/service/userService"

	"github.com/joho/godotenv"
)

func init()  {
	err := godotenv.Load()
	if err != nil {
	  	log.Fatal(err)
	}
}

func main()  {
	app.StartDB()

	db := app.GetDB()

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(db, userRepository)
	userController := userController.NewUserController(userService)

	categoryRepository := categoryrepository.NewCategoryRepository(db)
	categoryService := categoryservice.NewCategoryService(db, categoryRepository, userRepository)
	categoryController := categorycontroller.NewCategoryController(categoryService)

	productRepository := productrepository.NewProductRepository(db)
	productService := productservice.NewProductService(db, productRepository, userRepository)
	productController := productcontroller.NewProductController(productService)

	transactionRepository := transactionrepository.NewTransactionRepository(db)
	transactionService := transactionservice.NewTransactionService(db, transactionRepository, userRepository, productRepository)
	transactionController := transactioncontroller.NewTransactionController(transactionService)

	app.Route(userController, categoryController, productController, transactionController)

}