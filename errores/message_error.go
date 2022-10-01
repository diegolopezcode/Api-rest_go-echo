package errores

import "errors"

var (
	ErrJson = errors.New("Error al convertir a json")
)
