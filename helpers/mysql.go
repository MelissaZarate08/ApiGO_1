package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/joho/godotenv"
)

type MySQLConnection struct {
	Conn *client.Conn
}


func NewMySQLConnection() (*MySQLConnection, error) {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("Error convirtiendo el puerto a entero: %v", err)
	}

	// Crear la dirección del servidor
	addr := fmt.Sprintf("%s:%d", host, port)

	
	conn, err := client.Connect(addr, username, password, database)
	if err != nil {
		return nil, fmt.Errorf("Error conectando a MySQL: %v", err)
	}

	fmt.Println("Conexión exitosa a MySQL")
	return &MySQLConnection{Conn: conn}, nil
}


func (conn *MySQLConnection) Close() {
	if conn.Conn != nil {
		conn.Conn.Close()
		fmt.Println("Conexión cerrada correctamente")
	}
}