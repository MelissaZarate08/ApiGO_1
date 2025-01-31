package domain

type Category struct {
	id   int32
	name string
}

func NewCategory(name string) *Category {
	return &Category{id: 1, name: name}
}

// ToJSON(): Este método convierte un objeto Category a un formato serializable en JSON
func (c *Category) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":   c.id,
		"name": c.name,
	}
}

// Métodos para acceder y modificar los atributos privados
func (c *Category) GetName() string {
	return c.name
}

func (c *Category) SetName(name string) {
	c.name = name
}

func (c *Category) GetID() int32 {
	return c.id
}

func (c *Category) SetID(id int32) {
	c.id = id
}
