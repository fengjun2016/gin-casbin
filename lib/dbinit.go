package lib

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Gorm *gorm.DB
)

func initDB() {
	Gorm = gormDB()
}

func gormDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mysqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	return db
}
