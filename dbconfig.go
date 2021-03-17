package main

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConfig() *gorm.DB {
	databaseName := os.Getenv("DATABASE_NAME")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePass := os.Getenv("DATABASE_PASS")
	databaseUrl := os.Getenv("DATABASE_URL")
	dsn := databaseUsername + ":" + databasePass + "@tcp(" + databaseUrl + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	log.Printf("DB Connected")
	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}
