package controllers

import (
	"net/http"
	"strconv"
	"holamundo/product/application"

	"github.com/gin-gonic/gin"
)

type GetProductByIDController struct {
	GetProductByIDUseCase *application.GetProductByID
}

func NewGetProductByIDController(getByID *application.GetProductByID) *GetProductByIDController {
	return &GetProductByIDController{GetProductByIDUseCase: getByID}
}

func (ctrl *GetProductByIDController) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product, err := ctrl.GetProductByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product.ToJSON()})
}
