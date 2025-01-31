package domain

type IProduct interface {
	Create(product Product) (int32, error)
	GetAll() ([]Product, error)
	GetByID(id int32) (Product, error)
	Update(product Product) error
	Delete(id int32) error
}
