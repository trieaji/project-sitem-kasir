package app

import (
	"projectsistemkasir/controller"
	"projectsistemkasir/middleware"

	"github.com/julienschmidt/httprouter"
)

func RouterAuth(categoryController controller.CategoryController, paymentController controller.PaymentController, userController controller.UserController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	//register
	router.POST("/api/users", userController.Create)

	//login
	router.POST("/api/auth/users/login", userController.Login)

	//categories
	router.POST("/api/categories",middleware.AuthJWTMiddleware(categoryController.Create))
	router.PUT("/api/categories/:categoryId",middleware.AuthJWTMiddleware(categoryController.Update))
	router.DELETE("/api/categories/:categoryId",middleware.AuthJWTMiddleware(categoryController.Delete))
	router.GET("/api/categories/:categoryId",middleware.AuthJWTMiddleware(categoryController.FindById))
	router.GET("/api/categories",middleware.AuthJWTMiddleware(categoryController.FindAll))

	//payments
	router.POST("/api/payments",middleware.AuthJWTMiddleware(paymentController.Create))
	router.PUT("/api/payments/:paymentId",middleware.AuthJWTMiddleware(paymentController.Update))
	router.DELETE("/api/payments/:paymentId",middleware.AuthJWTMiddleware(paymentController.Delete))
	router.GET("/api/payments/:paymentId",middleware.AuthJWTMiddleware(paymentController.FindById))
	router.GET("/api/payments",middleware.AuthJWTMiddleware(paymentController.FindAll))

	//users
	router.PUT("/api/users/:userId",middleware.AuthJWTMiddleware(userController.Update))
	router.DELETE("/api/users/:userId",middleware.AuthJWTMiddleware(userController.Delete))
	router.GET("/api/users/:userId",middleware.AuthJWTMiddleware(userController.FindById))
	router.GET("/api/users",middleware.AuthJWTMiddleware(userController.FindAll))
	router.DELETE("/api/auth/users/logout",middleware.AuthJWTMiddleware(userController.Logout))

	//product
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId",middleware.AuthJWTMiddleware(productController.Update))
	router.DELETE("/api/products/:productId",productController.Delete)
	router.GET("/api/products/:productId",middleware.AuthJWTMiddleware(productController.FindById))
	router.GET("/api/products",productController.FindAll)

	return router
}