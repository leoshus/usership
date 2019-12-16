package db

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"usership/config"
)

var db *pg.DB
//initialize datatbase
func Init() {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     config.Get().DB.Host + ":" + fmt.Sprint(config.Get().DB.Port),
			User:     config.Get().DB.User,
			Password: config.Get().DB.Password,
			Database: config.Get().DB.Database,
		})
	}
}

//get databse connection
func DB() *pg.DB {
	return db
}


