package main

import (
	"net/http"
	"projectsistemkasir/app"
	"projectsistemkasir/controller"
	"projectsistemkasir/helper"
	// "projectsistemkasir/middleware"
	"projectsistemkasir/repository"
	"projectsistemkasir/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/julienschmidt/httprouter"
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

	router := app.RouterAuth(categoryController,paymentController,userController)

	// router := httprouter.New()

	//categories
	// router.POST("/api/categories", categoryController.Create)
	// router.PUT("/api/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	// router.GET("/api/categories/:categoryId", categoryController.FindById)
	// router.GET("/api/categories", categoryController.FindAll)

	//payments
	// router.POST("/api/payments", paymentController.Create)
	// router.PUT("/api/payments/:paymentId", paymentController.Update)
	// router.DELETE("/api/payments/:paymentId", paymentController.Delete)
	// router.GET("/api/payments/:paymentId", paymentController.FindById)
	// router.GET("/api/payments", paymentController.FindAll)

	//users
	// router.POST("/api/users", userController.Create)
	// router.PUT("/api/users/:userId", userController.Update)
	// router.DELETE("/api/users/:userId", userController.Delete)
	// router.GET("/api/users/:userId", userController.FindById)
	// router.GET("/api/users", userController.FindAll)

	//login
	// router.POST("/api/auth/users/login", userController.Login)

	//logout
	// router.DELETE("/api/auth/users/logout", userController.Logout)

	// router.POST("/api/examples", exampleController.Create)

	server := http.Server {
		Addr:"localhost:4007",
		// Handler:middleware.NewAuthMiddleware(router),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}