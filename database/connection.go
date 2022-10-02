package database

import (
	"fmt"
	"sync"

	"github.com/DiegoLopez-ing/api_rest/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ones sync.Once
var gdb *gorm.DB

func GetDBConnection(config *configs.DbConfigs) (*gorm.DB, error) {
	var errr error
	fmt.Println(config)
	ones.Do(func() {
		schema := schema.NamingStrategy{SingularTable: true}
		db, err := gorm.Open(postgres.Open(config.GetConnectionString()), &gorm.Config{NamingStrategy: schema})
		errr = err
		gdb = db
	})
	return gdb, errr
}
