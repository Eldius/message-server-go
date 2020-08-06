package repository

import (
	"log"

	"github.com/Eldius/auth-server-go/config"
	"github.com/Eldius/auth-server-go/user"
	"github.com/jinzhu/gorm"

	// I need this
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
		panic("failed to connect database")
	}
	db.AutoMigrate(&user.CredentialInfo{}, &user.Profile{})
	return db
}

// SaveUser saves the new user credential
func SaveUser(c *user.CredentialInfo) {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Save(c).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit
		return nil
	})
	if err != nil {
		log.Panicln("Failed to insert data\n", err.Error())
	}
}
