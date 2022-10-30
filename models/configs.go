package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = GetDB()
	db.AutoMigrate(&User{}, &Category{}, &Article{},
		&Comment{})
}

func GetDB() *gorm.DB {
	dsn := "root:rafi213fajri@/bloggo?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
