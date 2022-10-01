package router

import (
	"github.com/labstack/echo/v4"
)

func UpgradeUsuario(e *echo.Echo) {
	e.Group(paths.USUARIO)

}
