package repository

import (
	"log"

	"github.com/Eldius/message-server-go/config"
	"github.com/Eldius/message-server-go/messenger"
	"github.com/jinzhu/gorm"

	// I need this
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// GetDB gets a database connection
func GetDB() *gorm.DB {
	if db == nil {
		db = initDB()
	}
	return db
}

func initDB() *gorm.DB {
	db, err := gorm.Open(config.GetDBEngine(), config.GetDBURL())
	if err != nil {
		log.Printf("failed to connect database to app database\n- engine: %s\n- url: %s\n", config.GetDBEngine(), config.GetDBURL())
		panic(err.Error())
	}
	if config.GetDBLogQueries() {
		db.LogMode(true)
	}
	db.AutoMigrate(&messenger.Message{})
	return db
}
