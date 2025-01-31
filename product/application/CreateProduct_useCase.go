package application

import "holamundo/product/domain"

type CreateProduct struct {
	repo domain.IProduct
}

func NewCreateProduct(repo domain.IProduct) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (uc *CreateProduct) Run(product *domain.Product) error {
	id, err := uc.repo.Create(*product)
	if err != nil {
		return err
	}
	product.SetID(id)
	return nil
}
