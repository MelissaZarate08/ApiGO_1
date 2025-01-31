package application
import "holamundo/product/domain"

type GetAllProduct struct {
	db domain.IProduct
}

func NewGetAllProduct(db domain.IProduct) *GetAllProduct {
	return &GetAllProduct{db: db}
}

func (uc *GetAllProduct) Run() ([]map[string]interface{}, error) {
	products, err := uc.db.GetAll()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, product := range products {
		result = append(result, product.ToJSON())
	}
	return result, nil
}