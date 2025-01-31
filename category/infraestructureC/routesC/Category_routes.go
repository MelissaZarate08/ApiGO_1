package routesC

import (
	"holamundo/category/infraestructureC/controllers"
	"holamundo/category/infraestructureC"
	"holamundo/category/application"

	"github.com/gin-gonic/gin"
)

//registra las rutas HTTP y conecta los controladores correspondientes

func RegisterCategoryRoutes(r *gin.Engine, mysqlC *infraestructureC.MySQLCategory) {
	// Controladores de los métodos
	controllerCreateC := controllers.NewCreateCategoryController(
		application.NewCreateCategory(mysqlC),
	)

	controllerGetC := controllers.NewGetAllCategoryController(
		application.NewGetAllCategories(mysqlC),
	)

	controllerGetByIDC := controllers.NewGetCategoryByIDController(
		application.NewGetCategoryByID(mysqlC),
	)
	controllerUpdateC := controllers.NewUpdateCategoryController(
		application.NewUpdateCategory(mysqlC),
	)
	controllerDeleteC := controllers.NewDeleteCategoryController(
		application.NewDeleteCategory(mysqlC),
	)

	r.POST("/categories", controllerCreateC.CreateCategory)
	r.GET("/categories", controllerGetC.GetAllCategory)
	r.GET("/categories/:id", controllerGetByIDC.GetCategoryByID)
	r.PUT("/categories/:id", controllerUpdateC.UpdateCategory)
	r.DELETE("/categories/:id", controllerDeleteC.DeleteCategory)
}
