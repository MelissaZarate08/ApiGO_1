package application
import "holamundo/product/domain"

type GetProductByID struct {
	db domain.IProduct
}

func NewGetProductByID(db domain.IProduct) *GetProductByID {
	return &GetProductByID{db: db}
}

func (uc *GetProductByID) Run(id int32) (domain.Product, error) {
    return uc.db.GetByID(id)
}