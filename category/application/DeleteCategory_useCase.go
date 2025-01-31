package application

import "holamundo/category/domain"

type DeleteCategory struct {
	repository domain.ICategory
}

func NewDeleteCategory(repository domain.ICategory) *DeleteCategory {
	return &DeleteCategory{repository}
}

func (d *DeleteCategory) Run(id int32) error {
	return d.repository.Delete(id)
}
