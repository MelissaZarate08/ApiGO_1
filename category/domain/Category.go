package domain

type Category struct {
	id     int32
	name   string
	status int32
}

// Nueva función constructora que reciba también el status
func NewCategory(name string, status int32) *Category {
	return &Category{
		name:   name,
		status: status,
	}
}

// ToJSON: convierte un objeto Category a un formato serializable en JSON
func (c *Category) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":     c.id,
		"name":   c.name,
		"status": c.status,
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

func (c *Category) GetStatus() int32 {
	return c.status
}

func (c *Category) SetStatus(status int32) {
	c.status = status
}
