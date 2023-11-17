package models

import (
	"github.com/nanoohlaing1997/online-ordering-items/env"
	"github.com/nanoohlaing1997/online-ordering-items/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var environ = env.GetEnviroment()

type DatabaseManger struct {
	*ItemDB
	*CategoryDB
	*OrderDB
	*UserDB
}

func DBConn(conn string) *gorm.DB {
	conn = conn + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: log.GormLog(),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func NewDatabaseManager() *DatabaseManger {
	return &DatabaseManger{
		ItemManager(DBConn(environ.DbConfig.OrderDB)),
		CategoryManager(DBConn(environ.DbConfig.OrderDB)),
		OrderManager(DBConn(environ.DbConfig.OrderDB)),
		UserManager(DBConn(environ.DbConfig.OrderDB)),
	}
}
