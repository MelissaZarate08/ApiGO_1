package controllers

import (
	"net/http"
	"strconv"
	"holamundo/category/application"
	"holamundo/category/domain"

	"github.com/gin-gonic/gin"
)

type UpdateCategoryController struct {
	UpdateCategoryUseCase *application.UpdateCategory
}

func NewUpdateCategoryController(update *application.UpdateCategory) *UpdateCategoryController {
	return &UpdateCategoryController{UpdateCategoryUseCase: update}
}

func (ctrl *UpdateCategoryController) UpdateCategory(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Status int32  `json:"status"`
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

	// Usamos el constructor y luego seteamos el ID
	category := domain.NewCategory(req.Name, req.Status)
	category.SetID(int32(id))

	err = ctrl.UpdateCategoryUseCase.Run(*category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoría actualizada"})
}
