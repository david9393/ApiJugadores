package handlers

import (
	"github.com/david9393/ApiJugadores/routers"
	"github.com/labstack/echo"
)

func RouteApiFifa(e *echo.Echo) {
	menu := e.Group("/fifa")
	menu.POST("/add", routers.ObtenerJugadores)
}
