package application

import "holamundo/category/domain"

type CreateCategory struct {
	repo domain.ICategory
}

func NewCreateCategory(repo domain.ICategory) *CreateCategory {
	return &CreateCategory{repo: repo}
}

func (uc *CreateCategory) Run(category *domain.Category) error {
	id, err := uc.repo.Create(*category)
	if err != nil {
		return err
	}
	category.SetID(id)
	return nil
}
