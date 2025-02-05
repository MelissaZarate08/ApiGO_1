package infraestructureC

import (
	"fmt"
	"holamundo/category/domain"

	"github.com/go-mysql-org/go-mysql/client"
)

type MySQLCategory struct {
	Conn *client.Conn
}

// Create: Inserta la categoría con el campo status
func (mysqlC *MySQLCategory) Create(category domain.Category) (int32, error) {
	query := "INSERT INTO categories (name, status) VALUES (?, ?)"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(category.GetName(), category.GetStatus())
	if err != nil {
		return 0, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	lastInsertId := int32(result.InsertId)
	return lastInsertId, nil
}

// GetAll: Ahora seleccionamos también el campo status
func (mysqlC *MySQLCategory) GetAll() ([]domain.Category, error) {
	query := "SELECT id, name, status FROM categories"
	rows, err := mysqlC.Conn.Execute(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo categorías: %v", err)
	}

	var categories []domain.Category
	for _, row := range rows.Values {
		id := row[0].AsInt64()
		name := string(row[1].AsString()) // Conversión a string
		status := row[2].AsInt64()

		category := domain.NewCategory(name, int32(status))
		category.SetID(int32(id))
		categories = append(categories, *category)
	}

	return categories, nil
}

// GetByID: también leemos el campo status
func (mysqlC *MySQLCategory) GetByID(id int32) (domain.Category, error) {
	query := "SELECT id, name, status FROM categories WHERE id = ?"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return domain.Category{}, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(id)
	if err != nil {
		return domain.Category{}, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	if len(result.Values) == 0 {
		return domain.Category{}, fmt.Errorf("categoría con ID %d no encontrada", id)
	}

	row := result.Values[0]
	idFromDB := row[0].AsInt64()
	name := string(row[1].AsString()) // Conversión a string
	status := row[2].AsInt64()

	category := domain.NewCategory(name, int32(status))
	category.SetID(int32(idFromDB))
	return *category, nil
}

// Update: actualizamos también el campo status
func (mysqlC *MySQLCategory) Update(category domain.Category) error {
	query := "UPDATE categories SET name = ?, status = ? WHERE id = ?"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(category.GetName(), category.GetStatus(), category.GetID())
	if err != nil {
		return fmt.Errorf("error actualizando categoría: %v", err)
	}

	if result.AffectedRows == 0 {
		return fmt.Errorf("categoría con ID %d no encontrada", category.GetID())
	}

	return nil
}

// Delete: permanece igual
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

// Nuevo método: obtener categorías filtradas por status
func (mysqlC *MySQLCategory) GetByStatus(status int32) ([]domain.Category, error) {
	query := "SELECT id, name, status FROM categories WHERE status = ?"
	stmt, err := mysqlC.Conn.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(status)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	var categories []domain.Category
	for _, row := range result.Values {
		idFromDB := row[0].AsInt64()
		name := string(row[1].AsString()) // Conversión a string
		statusDB := row[2].AsInt64()

		category := domain.NewCategory(name, int32(statusDB))
		category.SetID(int32(idFromDB))
		categories = append(categories, *category)
	}

	return categories, nil
}
