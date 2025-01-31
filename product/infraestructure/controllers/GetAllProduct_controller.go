package controllers

import (
	"net/http"
	"holamundo/product/application"

	"github.com/gin-gonic/gin"
)

type GetAllProductController struct {
	GetAllProductsUseCase *application.GetAllProduct
}

func NewGetAllProductController(getAll *application.GetAllProduct) *GetAllProductController {
	return &GetAllProductController{GetAllProductsUseCase: getAll}
}

func (ctrl *GetAllProductController) GetAllProducts(c *gin.Context) {
	products, err := ctrl.GetAllProductsUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay productos disponibles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}