package application

import "holamundo/product/domain"

type UpdateProduct struct {
	repository domain.IProduct
}

func NewUpdateProduct(repository domain.IProduct) *UpdateProduct {
	return &UpdateProduct{repository}
}

func (u *UpdateProduct) Run(product domain.Product) error {
	return u.repository.Update(product)
}