package repository

import (
	"log"
	"time"

	"github.com/eldius/message-server-go/config"
	"github.com/eldius/message-server-go/logger"
	"github.com/eldius/message-server-go/messenger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
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
	newLogger := gormlogger.New(
		logger.Logger().WithFields(logrus.Fields{
			"db_engine": config.GetDBEngine(),
		}),
		gormlogger.Config{
			SlowThreshold:             1 * time.Second, // Slow SQL threshold
			LogLevel:                  gormlogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,           // Disable color
		},
	)
	db, err := gorm.Open(GetDialect(), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("failed to connect database to app database\n- engine: %s\n- url: %s\n", config.GetDBEngine(), config.GetDBURL())
		panic(err.Error())
	}
	if config.GetDBLogQueries() {
		//db.LogMode
	}
	db.AutoMigrate(&messenger.Message{})
	return db
}

func GetDialect() gorm.Dialector {
	switch config.GetDBEngine() {
	case "sqlite":
		return sqlite.Open(config.GetDBURL())
	case "mysql", "mariadb":
		return mysql.Open(config.GetDBURL())
	default:
		return sqlite.Open(config.GetDBURL())
	}
}
