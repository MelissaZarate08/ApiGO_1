package controllers

import (
	"net/http"
	"strconv"
	"holamundo/category/application"

	"github.com/gin-gonic/gin"
)

type GetCategoriesByStatusController struct {
	GetCategoriesByStatusUseCase *application.GetCategoriesByStatus
}

func NewGetCategoriesByStatusController(usecase *application.GetCategoriesByStatus) *GetCategoriesByStatusController {
	return &GetCategoriesByStatusController{usecase}
}

func (ctrl *GetCategoriesByStatusController) GetCategoriesByStatus(c *gin.Context) {
	statusParam := c.Param("status")
	statusInt, err := strconv.Atoi(statusParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}

	categories, err := ctrl.GetCategoriesByStatusUseCase.Run(int32(statusInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(categories) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay categorías para este status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}
