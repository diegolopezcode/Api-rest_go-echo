package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ones sync.once
var gdb *gorm.DB

func GetDBConnection(config *configs.DbConfig) (*gorm.DB, error) {
	var errr error
	ones.Do(func() {
		schema := &schema.NamingStrategy{
			SingularTable: true,
		}
		db, err = gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{
			NamingStrategy: schema,
		})
		errr = err
		gdb = db
	})
	return db, err
}
