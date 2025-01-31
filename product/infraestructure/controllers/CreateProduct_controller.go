package controllers

import (
	"net/http"
	"holamundo/product/application"
	"holamundo/product/domain"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	CreateProductUseCase *application.CreateProduct
}

func NewCreateProductController(create *application.CreateProduct) *CreateProductController {
	return &CreateProductController{CreateProductUseCase: create}
}

func (ctrl *CreateProductController) CreateProduct(c *gin.Context) {
	var req struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	product := domain.NewProduct(req.Name, req.Price)
	err := ctrl.CreateProductUseCase.Run(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Producto creado", "product": product.ToJSON()})
}