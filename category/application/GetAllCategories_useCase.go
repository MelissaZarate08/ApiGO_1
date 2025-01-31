package application

import "holamundo/category/domain"

type GetAllCategories struct {
	db domain.ICategory
}

func NewGetAllCategories(db domain.ICategory) *GetAllCategories {
	return &GetAllCategories{db: db}
}

func (uc *GetAllCategories) Run() ([]map[string]interface{}, error) {
	categories, err := uc.db.GetAll()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, category := range categories {
		result = append(result, category.ToJSON())
	}
	return result, nil
}
