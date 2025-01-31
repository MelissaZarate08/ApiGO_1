package routes

import (
	"holamundo/product/infraestructure/controllers"
	"holamundo/product/infraestructure"
	"holamundo/product/application"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, mysql *infraestructure.MySQL) {
	// Crear controlador de producto
	controllerCreateP := controllers.NewCreateProductController(
		application.NewCreateProduct(mysql),
	)

	controllerGetP := controllers.NewGetAllProductController(
		application.NewGetAllProduct(mysql),
	)

	controllerGetByIDP := controllers.NewGetProductByIDController(
		application.NewGetProductByID(mysql),
	)
	controllerUpdateP := controllers.NewUpdateProductController(
		application.NewUpdateProduct(mysql),
	)
	controllerDeleteP := controllers.NewDeleteProductController(
		application.NewDeleteProduct(mysql),
	)

	r.POST("/products", controllerCreateP.CreateProduct)
	r.GET("/products", controllerGetP.GetAllProducts)
	r.GET("/products/:id", controllerGetByIDP.GetProductByID)
	r.PUT("/products/:id", controllerUpdateP.UpdateProduct)
	r.DELETE("/products/:id", controllerDeleteP.DeleteProduct)
}