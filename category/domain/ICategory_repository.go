package domain

type ICategory interface {
	Create(category Category) (int32, error)
	GetAll() ([]Category, error)
	GetByID(id int32) (Category, error)
	Update(category Category) error
	Delete(id int32) error
}
