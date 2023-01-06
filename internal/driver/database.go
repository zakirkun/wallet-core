package driver

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBModel struct {
	Server     string
	DBname     string
	DBUsername string
	DBPassword string
	DBPort     int
}

type IDBInstance interface {
	Open() *gorm.DB
}

func NewDatabase(opt DBModel) IDBInstance {
	return DBModel{opt.Server, opt.DBname, opt.DBUsername, opt.DBPassword, opt.DBPort}
}

func (i DBModel) Open() *gorm.DB {

	var DSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", i.DBUsername, i.DBUsername, i.Server, i.DBPort, i.DBname)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
