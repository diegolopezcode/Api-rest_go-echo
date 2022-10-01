package handlers

import (
	"github.com/Diegolopez-ing/api_rest/errores"
	"github.com/Diegolopez-ing/api_rest/models"
	"github.com/Diegolopez-ing/api_rest/services"
	"github.com/Diegolopez-ing/api_rest/utils"
	"github.com/labstack/echo"
)

type UsuarioHandler struct {
	binder  echo.DefaultBinder
	service *services.UsuarioService
}

func NewUsuarioHandler(service *services.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{
		binder:  echo.DefaultBinder{},
		service: service,
	}
}

func (handler *UsuarioHandler) LogginHandler(c echo.Context) error {
	rs := models.RequestSession{}
	if err := handler.binder.Bind(&rs, c); err != nil {
		return errores.GenErrResposeEcho(err, c)
	}
	usuario, err := handler.service.Loggin(rs)
	if err != nil {
		return errores.GenErrResposeEcho(err, c)
	}
	return utils.OkResponse(c, usuario)
}
