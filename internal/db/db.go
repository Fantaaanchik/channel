package db

import (
	"chans/config"
	"chans/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func ConnectionToDB() *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(config.Configure.DB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db, err: ", err.Error())

	}
	err = db.AutoMigrate(&models.Card{})
	return db
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	err = sqlDB.Close()
	if err != nil {
		return
	}

}
