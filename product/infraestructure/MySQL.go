package infraestructure

import (
	"fmt"
	"holamundo/product/domain"

	"github.com/go-mysql-org/go-mysql/client"
)


type MySQL struct {
	Conn *client.Conn
}


func (mysql *MySQL) Create(product domain.Product) (int32, error) {
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	stmt, err := mysql.Conn.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(product.GetName(), product.GetPrice())
	if err != nil {
		return 0, fmt.Errorf("error ejecutando consulta: %v", err)
	}

	lastInsertId := int32(result.InsertId)
	return lastInsertId, nil
}

func (mysql *MySQL) GetAll() ([]domain.Product, error) {
    query := "SELECT id, name, price FROM products"
    rows, err := mysql.Conn.Execute(query)
    if err != nil {
        return nil, fmt.Errorf("error obteniendo productos: %v", err)
    }

    var products []domain.Product

    fmt.Printf("Número de filas obtenidas: %d\n", len(rows.Values))

    for _, row := range rows.Values {
        id := row[0].AsInt64()
        name := string(row[1].AsString())
        price := float32(row[2].AsFloat64())
        fmt.Printf("Producto: ID=%d, Name=%s, Price=%.2f\n", id, name, price)

        product := domain.NewProduct(name, price)
        product.SetID(int32(id))
        products = append(products, *product)
    }

    if len(products) == 0 {
        fmt.Println("No se encontraron productos")
    }

    return products, nil
}

func (mysql *MySQL) GetByID(id int32) (domain.Product, error) {
    query := "SELECT id, name, price FROM products WHERE id = ?"
    stmt, err := mysql.Conn.Prepare(query)
    if err != nil {
        return domain.Product{}, fmt.Errorf("error preparando consulta: %v", err)
    }
    defer stmt.Close()

    result, err := stmt.Execute(id)
    if err != nil {
        return domain.Product{}, fmt.Errorf("error ejecutando consulta: %v", err)
    }

    fmt.Printf("Filas obtenidas para ID=%d: %d\n", id, len(result.Values))

    if len(result.Values) == 0 {
        return domain.Product{}, fmt.Errorf("producto con ID %d no encontrado", id)
    }

    row := result.Values[0]
    idFromDB := row[0].AsInt64()
    name := string(row[1].AsString())
    price := float32(row[2].AsFloat64())

    fmt.Printf("Producto encontrado: ID=%d, Name=%s, Price=%.2f\n", idFromDB, name, price)

	// Crea una instancia del producto y lo devuelve.
    product := domain.NewProduct(name, price)
    product.SetID(int32(idFromDB))
    return *product, nil
}

func (mysql *MySQL) Update(product domain.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	stmt, err := mysql.Conn.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	
	result, err := stmt.Execute(product.GetName(), product.GetPrice(), product.GetID())
	if err != nil {
		return fmt.Errorf("error actualizando producto: %v", err)
	}

	if result.AffectedRows == 0 {
		return fmt.Errorf("producto con ID %d no encontrado", product.GetID())
	}

	return nil
}


func (mysql *MySQL) Delete(id int32) error {
	query := "DELETE FROM products WHERE id = ?"
	stmt, err := mysql.Conn.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparando consulta: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Execute(id)
	if err != nil {
		return fmt.Errorf("error eliminando producto: %v", err)
	}

	if result.AffectedRows == 0 {
		return fmt.Errorf("producto con ID %d no encontrado", id)
	}

	return nil
}