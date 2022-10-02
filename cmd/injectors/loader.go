package injectors

import (
	"sync"

	"github.com/DiegoLopez-ing/api_rest/configs"
	"github.com/DiegoLopez-ing/api_rest/database"
<<<<<<< HEAD
=======
	"github.com/DiegoLopez-ing/api_rest/handlers"
	"github.com/DiegoLopez-ing/api_rest/repos"
	"github.com/DiegoLopez-ing/api_rest/services"
>>>>>>> 34b956835237401f626c95762799cba11afd2588
	"gorm.io/gorm"
)

var ones sync.Once

func LoadConfig(configFilepath string) error {
	var errr error
	ones.Do(func() {
		if conf, err := configs.LoadDbConfig(configFilepath); err != nil {
			errr = err
			return
		} else {
			con, er := database.GetDBConnection(conf)
			if err != nil {
				errr = er

				return
			}
			initRepository(con)
			initServices()
			initHandlers()
		}
	})
	return errr
}

func initRepository(con *gorm.DB) {
	usuarioRepository = repos.NewUsuarioRepository(con)
}

func initServices() {
	usuarioService = services.NewUsuarioService(usuarioRepository)
}

func initHandlers() {
	usuarioHandler = handlers.NewUsuarioHandler(usuarioService)
}
