package controllers

import (
	"net/http"
	"holamundo/category/application"

	"github.com/gin-gonic/gin"
)

type GetAllCategoryController struct {
	GetAllCategoriesUseCase *application.GetAllCategories
}

func NewGetAllCategoryController(getAll *application.GetAllCategories) *GetAllCategoryController {
	return &GetAllCategoryController{GetAllCategoriesUseCase: getAll}
}

func (ctrl *GetAllCategoryController) GetAllCategory(c *gin.Context) {
	categories, err := ctrl.GetAllCategoriesUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(categories) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay categorías disponibles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}