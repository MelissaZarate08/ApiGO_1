package application

import "holamundo/category/domain"

type UpdateCategory struct {
	repository domain.ICategory
}

func NewUpdateCategory(repository domain.ICategory) *UpdateCategory {
	return &UpdateCategory{repository}
}

func (u *UpdateCategory) Run(category domain.Category) error {
	return u.repository.Update(category)
}
