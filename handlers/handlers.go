package handlers

import (
	"github.com/david9393/ApiJugadores/routers"
	"github.com/labstack/echo"
)

func RouteApiFifa(e *echo.Echo) {
	menu := e.Group("api/v1")
	menu.POST("/add", routers.AddJugadores)
	menu.POST("/team", routers.GetJugadoresEquipo)
	menu.POST("/players", routers.GetJugadoresNombre)
}
