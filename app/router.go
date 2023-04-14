package app

import (
	"Dzaakk/go-restful-api/controller"
	"Dzaakk/go-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoriesId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoriesId", categoryController.Update)
	router.DELETE("/api/categories/:categoriesId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
