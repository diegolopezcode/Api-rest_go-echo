package router

import (
	"github.com/DiegoLopez-ing/api_rest/api/paths"
	"github.com/DiegoLopez-ing/api_rest/cmd/injectors"
	"github.com/labstack/echo/v4"
)

func UpgradeUsuario(e *echo.Echo) {
	h := injectors.GetUsuarioHandler()
	g := e.Group(paths.USUARIO)
	g.POST("/login", h.LogginHandler)
	g.GET("/", h.FindAll)

}
