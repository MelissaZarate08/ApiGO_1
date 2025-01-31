package controllers

import (
	"net/http"
	"strconv"
	"holamundo/product/application"
	"holamundo/product/domain"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	UpdateProductUseCase *application.UpdateProduct
}

func NewUpdateProductController(update *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{UpdateProductUseCase: update}
}

func (ctrl *UpdateProductController) UpdateProduct(c *gin.Context) {
	var req struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product := domain.NewProduct(req.Name, req.Price)
	product.SetID(int32(id))

	err = ctrl.UpdateProductUseCase.Run(*product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado"})
}