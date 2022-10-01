package injectors

import (
	"sync"

	"github.com/DiegoLopez-ing/api_rest/configs"
	"github.com/DiegoLopez-ing/api_rest/database"
	"gorm.io/gorm"
)

var ones sync.Once

func LoadConfig(configFilepath string) error {
	var errr error
	ones.Do(func() {
		if conf, err := configs.LoadDbConfigs(configFilepath); err != nil {
			errr = err
			return
		} else {
			con, er = database.GetDBConnextion(conf)
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
	Con = con
}

func initServices() {
}

func initHandlers() {

}
