package app

import (
	"projectsistemkasir/controller"
	// "projectsistemkasir/exception"

	"github.com/julienschmidt/httprouter"
)

func RouterAuth(categoryController controller.CategoryController, paymentController controller.PaymentController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	//categories
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)

	//payments
	router.POST("/api/payments", paymentController.Create)
	router.PUT("/api/payments/:paymentId", paymentController.Update)
	router.DELETE("/api/payments/:paymentId", paymentController.Delete)
	router.GET("/api/payments/:paymentId", paymentController.FindById)
	router.GET("/api/payments", paymentController.FindAll)

	//users
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)
	router.GET("/api/users/:userId", userController.FindById)
	router.GET("/api/users", userController.FindAll)
	router.DELETE("/api/auth/users/logout", userController.Logout)


	return router
}

func RouterWithoutAuth(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	//register
	router.POST("/api/users", userController.Create)
	//login
	router.POST("/api/auth/users/login", userController.Login)

	return router
}

// func mixRouter(paymentController controller.PaymentController) *httprouter.Router {
// 	router := httprouter.New()

// 	router.POST("/api/payments", paymentController.Create)

// 	// router.PanicHandler = exception.ErrorHandler

// 	return router
// }

