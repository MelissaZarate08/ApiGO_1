package infraestructureC

import (
	"fmt"
	"holamundo/category/domain"
	"github.com/go-mysql-org/go-mysql/client"
)

type MySQLCategory struct {
	Conn *client.Conn
}

//metodos
func (mysqlC *MySQLCategory) Create(category domain.Category) (int32, error) {
	query := "INSERT INTO categories (name) VALUES (?)"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(category.GetName())
	if err != nil {
		return 0, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	lastInsertId := int32(result.InsertId)
	return lastInsertId, nil
}

func (mysqlC *MySQLCategory) GetAll() ([]domain.Category, error) {
	query := "SELECT id, name FROM categories"
	rows, err := mysqlC.Conn.Execute(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo categorías: %v", err)
	}

	var categories []domain.Category

	fmt.Printf("Número de filas obtenidas: %d\n", len(rows.Values))

	for _, row := range rows.Values {
		id := row[0].AsInt64()
		name := string(row[1].AsString())
		fmt.Printf("Categoría: ID=%d, Name=%s\n", id, name)

		category := domain.NewCategory(name)
		category.SetID(int32(id))
		categories = append(categories, *category)
	}

	if len(categories) == 0 {
		fmt.Println("No se encontraron categorías")
	}

	return categories, nil
}

func (mysqlC *MySQLCategory) GetByID(id int32) (domain.Category, error) {
	query := "SELECT id, name FROM categories WHERE id = ?"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return domain.Category{}, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(id)
	if err != nil {
		return domain.Category{}, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	fmt.Printf("Filas obtenidas para ID=%d: %d\n", id, len(result.Values))

	if len(result.Values) == 0 {
		return domain.Category{}, fmt.Errorf("categoría con ID %d no encontrada", id)
	}

	row := result.Values[0]
	idFromDB := row[0].AsInt64()
	name := string(row[1].AsString())

	fmt.Printf("Categoría encontrada: ID=%d, Name=%s\n", idFromDB, name)

	category := domain.NewCategory(name)
	category.SetID(int32(idFromDB))
	return *category, nil
}

func (mysqlC *MySQLCategory) Update(category domain.Category) error {
	query := "UPDATE categories SET name = ? WHERE id = ?"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(category.GetName(), category.GetID())
	if err != nil {
		return fmt.Errorf("error actualizando categoría: %v", err)
	}

	if result.AffectedRows == 0 {
		return fmt.Errorf("categoría con ID %d no encontrada", category.GetID())
	}

	return nil
}

func (mysqlC *MySQLCategory) Delete(id int32) error {
	query := "DELETE FROM categories WHERE id = ?"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(id)
	if err != nil {
		return fmt.Errorf("error eliminando categoría: %v", err)
	}

	if result.AffectedRows == 0 {
		return fmt.Errorf("categoría con ID %d no encontrada", id)
	}

	return nil
}