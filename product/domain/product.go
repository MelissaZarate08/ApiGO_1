package domain

type Product struct {
	id int32
	name string
	price float32

}
func NewProduct(name string, price float32) *Product {
	return &Product{id:1,name: name, price: price}
}

// Método para devolver un producto serializable
func (p *Product) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id": p.id,
		"name":  p.name,
		"price": p.price,
	}
}

/*funcion para convertir un atributo privado a publico y se pueda utilizar*/
func (p *Product) GetName() string {
	return p.name
}

func (p * Product) SetName(name string){
	p.name =name
}

func (p *Product) GetPrice() float32 {
	return p.price
}

func (p *Product) SetPrice(price float32) {
	p.price = price
}

func (p *Product) GetID() int32 {
	return p.id
}

func (p *Product) SetID(id int32) {
	p.id = id
}