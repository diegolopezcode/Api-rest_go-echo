package handlers

import (
	"github.com/DiegoLopez-ing/api_rest/errores"
	"github.com/DiegoLopez-ing/api_rest/models"
	"github.com/DiegoLopez-ing/api_rest/services"
	"github.com/DiegoLopez-ing/api_rest/utils"
	"github.com/labstack/echo/v4"
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

func (handler *UsuarioHandler) FindAll(c echo.Context) error {

	return utils.OkResponse(c, "Todo estuvo Bien")
}
