package errores

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func logErrs(err error) {

	log.Println(err.Error())

}

func GenErrResposeEcho(err error, c echo.Context) error {
	go logErrs(err)
	code := http.StatusBadRequest
	message := ""
	if errors.Is(err, ErrJson) {
		message = "Error al convertir a json"
	}

	return c.JSON(code, message)
}
