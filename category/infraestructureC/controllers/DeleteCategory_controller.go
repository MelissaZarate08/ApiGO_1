package controllers

import (
	"net/http"
	"strconv"
	"holamundo/category/application"

	"github.com/gin-gonic/gin"
)

type DeleteCategoryController struct {
	DeleteCategoryUseCase *application.DeleteCategory
}

func NewDeleteCategoryController(delete *application.DeleteCategory) *DeleteCategoryController {
	return &DeleteCategoryController{DeleteCategoryUseCase: delete}
}

func (ctrl *DeleteCategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.DeleteCategoryUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoría eliminada"})
}