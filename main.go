package main

import (
	"log"

	"github.com/david9393/ApiJugadores/bd"
	"github.com/david9393/ApiJugadores/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	bd.ConectarBD()
	bd.MakeMigration()
	log.Println("Conexion exitosa")
	e := echo.New()
	e.Use(middleware.CORS())

	handlers.RouteApiFifa(e)
	log.Println("Servidor iniciado en el puerto 8080")

	err := e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
