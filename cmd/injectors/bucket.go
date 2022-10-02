package injectors

import (
	"github.com/DiegoLopez-ing/api_rest/handlers"
	"github.com/DiegoLopez-ing/api_rest/repos"
	"github.com/DiegoLopez-ing/api_rest/services"
	"gorm.io/gorm"
)

var (
	Connection *gorm.DB

	//REPOSITORY
	usuarioRepository *repos.UsuarioRepository

	//SERVICE
	usuarioService *services.UsuarioService

	//HANDLER
	usuarioHandler *handlers.UsuarioHandler
)
