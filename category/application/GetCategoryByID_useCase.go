package application

import "holamundo/category/domain"

type GetCategoryByID struct {
	db domain.ICategory
}

func NewGetCategoryByID(db domain.ICategory) *GetCategoryByID {
	return &GetCategoryByID{db: db}
}

func (uc *GetCategoryByID) Run(id int32) (domain.Category, error) {
    return uc.db.GetByID(id)
}
