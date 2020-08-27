package repository

import (
	"github.com/Eldius/auth-server-go/config"
	"github.com/Eldius/auth-server-go/logger"
	"github.com/Eldius/auth-server-go/user"
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
		panic("failed to connect database")
	}
	if config.GetDBLogQueries() {
		db.LogMode(true)
	}
	db.AutoMigrate(&user.CredentialInfo{}, &user.Profile{})
	return db
}

// SaveUser saves the new user credential
func SaveUser(c *user.CredentialInfo) {
	if c == nil {
		return
	}
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Save(c).Error; err != nil {
			// return any error will rollback
			logger.Logger().Println("Error saving credentials")
			logger.Logger().Println(err.Error())
			return err
		}
		// return nil will commit
		return nil
	})
	if err != nil {
		logger.Logger().Panicln("Failed to insert data\n", err.Error())
	}
}

// FindUser finds the user
func FindUser(username string) *user.CredentialInfo {

	u := user.CredentialInfo{}
	GetDB().Where("User = ?", username).First(&u)
	return &u
}

// ListUSers returns all users
func ListUSers() (r []user.CredentialInfo) {
	GetDB().Find(&r, "")
	return
}
