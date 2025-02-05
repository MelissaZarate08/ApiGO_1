package application

import "holamundo/category/domain"

type GetCategoriesByStatus struct {
	repository domain.ICategory
}

func NewGetCategoriesByStatus(repository domain.ICategory) *GetCategoriesByStatus {
	return &GetCategoriesByStatus{repository}
}

func (uc *GetCategoriesByStatus) Run(status int32) ([]map[string]interface{}, error) {
	categories, err := uc.repository.GetByStatus(status)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, category := range categories {
		result = append(result, category.ToJSON())
	}

	return result, nil
}
