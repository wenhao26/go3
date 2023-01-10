package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"coinsky_go_project/blog-apis/pkg/setting"
)

var db *gorm.DB

func init() {

	var err error

	dbType := setting.DBType
	dbName := setting.DBName
	user := setting.DBUser
	password := setting.DBPassword
	host := setting.DBHost
	tablePrefix := setting.DBTablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		log.Println("连接数据库异常：", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	db.Close()
}
