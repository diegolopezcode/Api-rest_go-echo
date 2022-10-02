package injectors

import "github.com/DiegoLopez-ing/api_rest/handlers"

func GetUsuarioHandler() *handlers.UsuarioHandler {
	return usuarioHandler
}
