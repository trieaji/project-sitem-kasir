package main

import (
	"net/http"
	"projectsistemkasir/app"
	"projectsistemkasir/controller"
	"projectsistemkasir/helper"
	"projectsistemkasir/repository"
	"projectsistemkasir/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	paymentRepository := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepository,db,validate)
	paymentController := controller.NewPaymentController(paymentService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db , validate)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db , validate)
	productController := controller.NewProductController(productService)

	router := app.RouterAuth(categoryController,paymentController,userController,productController)

	server := http.Server {
		Addr:"localhost:4007",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}