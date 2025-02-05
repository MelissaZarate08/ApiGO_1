package controllers

import (
	"net/http"
	"holamundo/category/application"
	"holamundo/category/domain"

	"github.com/gin-gonic/gin"
)

type CreateCategoryController struct {
	CreateCategoryUseCase *application.CreateCategory
}

func NewCreateCategoryController(create *application.CreateCategory) *CreateCategoryController {
	return &CreateCategoryController{CreateCategoryUseCase: create}
}

func (ctrl *CreateCategoryController) CreateCategory(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Status int32  `json:"status"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Ahora usamos el nuevo constructor con (name, status)
	category := domain.NewCategory(req.Name, req.Status)

	err := ctrl.CreateCategoryUseCase.Run(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Categoría creada",
		"category": category.ToJSON(),
	})
}
