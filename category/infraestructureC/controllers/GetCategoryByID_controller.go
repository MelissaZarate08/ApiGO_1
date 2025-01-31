package controllers

import (
	"net/http"
	"strconv"
	"holamundo/category/application"

	"github.com/gin-gonic/gin"
)

type GetCategoryByIDController struct {
	GetCategoryByIDUseCase *application.GetCategoryByID
}

func NewGetCategoryByIDController(getByID *application.GetCategoryByID) *GetCategoryByIDController {
	return &GetCategoryByIDController{GetCategoryByIDUseCase: getByID}
}

func (ctrl *GetCategoryByIDController) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	category, err := ctrl.GetCategoryByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category.ToJSON()})
}