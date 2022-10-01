package injectors

import "gorm.io/gorm"

var (
	Connection *gorm.DB

	//REPOSITORY
	usuarioRepository *repos.UsuarioRepository

	//SERVICE
	usuarioService *services.UsuarioService

	//HANDLER
	usuarioHandler *handlers.UsuarioHandler
)
