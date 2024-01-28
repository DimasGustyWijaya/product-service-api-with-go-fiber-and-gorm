package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnect() *gorm.DB {

	db, err := gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/product"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
