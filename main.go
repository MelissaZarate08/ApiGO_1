package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"holamundo/helpers"
	"holamundo/product/infraestructure"
	"holamundo/product/infraestructure/routes"
	"holamundo/category/infraestructureC"
	"holamundo/category/infraestructureC/routesC"
)

func main() {
	mysqlConn, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error al conectar a MySQL: %v", err)
	}
	defer mysqlConn.Close()

	// Se crean instancias de MySQL para productos y categorías, pasando la conexión a la base de datos.
	mysql := &infraestructure.MySQL{Conn: mysqlConn.Conn}
	mysqlC := &infraestructureC.MySQLCategory{Conn: mysqlConn.Conn}

	// Para manejar las rutas HTTP.
	r := gin.Default()

	// Registra las rutas para productos y categorías.
	routes.RegisterProductRoutes(r, mysql)
	routesC.RegisterCategoryRoutes(r, mysqlC)

	r.Run(":8080")
}
