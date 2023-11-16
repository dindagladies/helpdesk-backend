package config

import (
	"helpdesk/api/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	// db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/helpdeskdb?charset=utf8&parseTime=True&loc=Local")

	dsn := "root:@tcp(127.0.0.1:3306)/helpdeskdb?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(structs.User{})

	return db
}
